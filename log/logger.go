package stllog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"golang.org/x/exp/slices"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerror "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"
)

type Logger struct {
	config

	l   *log.Logger
	lvl *atomic.Uint32

	parent *Logger
	group  string
}

func Default(debug bool) *Logger {
	var logger *Logger
	if debug {
		logger = New(os.Stdout, LevelDebug)
	} else {
		logger = New(os.Stdout, LevelInfo)
	}
	logger.SetDefaultConfig(logger.WithDisplayLevel().WithDisplayTime().WithDisplayPosition().WithDisplayColor().WithDisplayGroup())
	return logger
}

func New(w io.Writer, level Level) *Logger {
	cfg := config{
		ctx:   context.Background(),
		level: level,
	}

	var lvl atomic.Uint32
	lvl.Store(uint32(level))

	l := log.New(w, "", log.Lmsgprefix)

	return &Logger{config: cfg, l: l, lvl: &lvl}
}
func (l *Logger) NewGroup(group string) *Logger {
	var lvl atomic.Uint32
	lvl.Store(uint32(l.GetLevel()))

	return &Logger{
		config: l.config,
		l:      log.New(l.l.Writer(), "", log.Lmsgprefix),
		lvl:    &lvl,
		parent: l,
		group:  group,
	}
}

func (l *Logger) SetDefaultConfig(cfg config) {
	l.config = cfg
}
func (l *Logger) GetLevel() Level {
	return Level(l.lvl.Load())
}
func (l *Logger) SetLevel(level Level) {
	l.lvl.Store(uint32(level))
}
func (l *Logger) GetGroup() string {
	return l.group
}

func (l *Logger) enable(lvl Level) bool {
	return lvl >= l.GetLevel()
}

func (l *Logger) log(msg string, cfgs ...config) error {
	cfg := stlslices.Last(cfgs)

	if !l.enable(cfg.level) {
		return nil
	}

	if cfg.displayGroup {
		var groups []string
		for cursorLog := l; cursorLog != nil; cursorLog = cursorLog.parent {
			if cursorLog.group != "" {
				groups = append(groups, cursorLog.group)
			}
		}
		if !stlslices.Empty(groups) {
			slices.Reverse(groups)
			msg = fmt.Sprintf("%s | %s", strings.Join(groups, " | "), msg)
		}
	}

	if cfg.displayPos {
		stack := stlos.GetCurrentCallStack(cfg.posSkip + 1)
		msg = fmt.Sprintf("%s:%d | %s", stack.File, stack.Line, msg)
	}

	if cfg.displayTimeFormat != "" {
		msg = fmt.Sprintf("%s | %s", time.Now().Format(cfg.displayTimeFormat), msg)
	}

	if cfg.displayStack {
		var stacks []runtime.Frame
		if stlslices.Empty(cfg.stacks) {
			stacks = stlos.GetCallStacks(100, cfg.posSkip+1)
		} else {
			stacks = cfg.stacks
		}
		stackStrs := stlslices.Map(stacks, func(_ int, stack runtime.Frame) string {
			return fmt.Sprintf("\t%s:%d", stack.File, stack.Line)
		})
		msg = fmt.Sprintf("%s\n%s", msg, strings.Join(stackStrs, "\n"))
	}

	if cfg.displayLevel {
		level := cfg.level.AlignString()
		if cfg.displayColor && writerIsTerminal(l.l.Writer()) {
			level = cfg.level.Style().Sprintf(level)
			msg = cfg.level.MsgColor().Sprintf("| %s", msg)
		}
		msg = level + msg
	}

	return l.l.Output(0, msg)
}

func (l *Logger) Output(msg string, cfgs ...config) error {
	return l.log(msg, stlslices.Last(cfgs, l.config).WithPositionSkip(1))
}
func (l *Logger) Print(a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.log(fmt.Sprint(a...), cfg.WithPositionSkip(1))
}
func (l *Logger) Println(a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.log(fmt.Sprintln(a...), cfg.WithPositionSkip(1))
}

func (l *Logger) commonOutput(level Level, a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.Output(fmt.Sprint(a...), cfg.WithLevel(level).WithPositionSkip(2))
}
func (l *Logger) commonOutputln(level Level, a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.Output(fmt.Sprintln(a...), cfg.WithLevel(level).WithPositionSkip(2))
}
func (l *Logger) commonOutputf(level Level, format string, a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.Output(fmt.Sprintf(format, a...), cfg.WithLevel(level).WithPositionSkip(2))
}
func (l *Logger) commonSmartOutput(level Level, a any, cfgs ...config) error {
	cfg := stlslices.Last(cfgs, l.config)
	cfg = cfg.WithLevel(level).WithPositionSkip(2)
	if level >= LevelPanic {
		cfg = cfg.WithDisplayStack()
	}
	if err, ok := a.(error); ok {
		cfg = cfg.WithDisplayStack()
		var stlerr stlerror.Error
		if errors.As(err, &stlerr) {
			cfg = cfg.WithStacks(stlerr.Stacks())
		}
	}
	return l.Output(fmt.Sprint(a), cfg)
}

func (l *Logger) Debug(a ...any) error   { return l.commonOutput(LevelDebug, a...) }
func (l *Logger) Debugln(a ...any) error { return l.commonOutputln(LevelDebug, a...) }
func (l *Logger) Debugf(format string, a ...any) error {
	return l.commonOutputf(LevelDebug, format, a...)
}
func (l *Logger) SmartDebug(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelDebug, a, cfgs...)
}

func (l *Logger) Trace(a ...any) error   { return l.commonOutput(LevelTrace, a...) }
func (l *Logger) Traceln(a ...any) error { return l.commonOutputln(LevelTrace, a...) }
func (l *Logger) Tracef(format string, a ...any) error {
	return l.commonOutputf(LevelTrace, format, a...)
}
func (l *Logger) SmartTrace(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelTrace, a, cfgs...)
}

func (l *Logger) Info(a ...any) error   { return l.commonOutput(LevelInfo, a...) }
func (l *Logger) Infoln(a ...any) error { return l.commonOutputln(LevelInfo, a...) }
func (l *Logger) Infof(format string, a ...any) error {
	return l.commonOutputf(LevelInfo, format, a...)
}
func (l *Logger) SmartInfo(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelInfo, a, cfgs...)
}

func (l *Logger) Warn(a ...any) error   { return l.commonOutput(LevelWarn, a...) }
func (l *Logger) Warnln(a ...any) error { return l.commonOutputln(LevelWarn, a...) }
func (l *Logger) Warnf(format string, a ...any) error {
	return l.commonOutputf(LevelWarn, format, a...)
}
func (l *Logger) SmartWarn(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelWarn, a, cfgs...)
}

func (l *Logger) Keyword(a ...any) error   { return l.commonOutput(LevelKeyword, a...) }
func (l *Logger) Keywordln(a ...any) error { return l.commonOutputln(LevelKeyword, a...) }
func (l *Logger) Keywordf(format string, a ...any) error {
	return l.commonOutputf(LevelKeyword, format, a...)
}
func (l *Logger) SmartKeyword(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelKeyword, a, cfgs...)
}

func (l *Logger) Error(a ...any) error   { return l.commonOutput(LevelError, a...) }
func (l *Logger) Errorln(a ...any) error { return l.commonOutputln(LevelError, a...) }
func (l *Logger) Errorf(format string, a ...any) error {
	return l.commonOutputf(LevelError, format, a...)
}
func (l *Logger) SmartError(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelError, a, cfgs...)
}

func (l *Logger) Panic(a ...any) error   { return l.commonOutput(LevelPanic, a...) }
func (l *Logger) Panicln(a ...any) error { return l.commonOutputln(LevelPanic, a...) }
func (l *Logger) Panicf(format string, a ...any) error {
	return l.commonOutputf(LevelPanic, format, a...)
}
func (l *Logger) SmartPanic(a any, cfgs ...config) error {
	return l.commonSmartOutput(LevelPanic, a, cfgs...)
}
