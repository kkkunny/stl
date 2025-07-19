package stllog

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlerr "github.com/kkkunny/stl/error"
	"github.com/kkkunny/stl/internal/slices"
	stlos "github.com/kkkunny/stl/os"
	stlval "github.com/kkkunny/stl/value"
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
func (l *Logger) SetWriter(w io.Writer) {
	l.l.SetOutput(w)
}
func (l *Logger) GetWriter() io.Writer {
	return l.l.Writer()
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
		frame := stlval.ValueOr(cfg.displayPosFrame, stlos.WrapRuntimeFrame(stlos.GetCurrentCallStack(cfg.posSkip+1)))
		msg = fmt.Sprintf("%s:%d | %s", frame.File(), frame.Line(), msg)
	}

	if cfg.displayTimeFormat != "" {
		msg = fmt.Sprintf("%s | %s", time.Now().Format(cfg.displayTimeFormat), msg)
	}

	if cfg.displayStack {
		var frames []stlos.Frame
		if stlslices.Empty(cfg.frames) {
			frames = stlslices.Map(stlos.GetCallStacks(32, cfg.posSkip+1), func(_ int, f runtime.Frame) stlos.Frame {
				return stlos.WrapRuntimeFrame(f)
			})
		} else {
			frames = cfg.frames
		}
		stackStrs := stlslices.Map(frames, func(_ int, f stlos.Frame) string {
			return fmt.Sprintf("\t%s:%d", f.File(), f.Line())
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
	if len(a) == 1 {
		cfg = cfg.WithLevel(level).WithPositionSkip(2)
		if level >= LevelPanic {
			cfg = cfg.WithDisplayStack()
		}
		if err, ok := a[0].(error); ok {
			cfg = cfg.WithDisplayStack()
			if frames := stlerr.GetErrorStackFrames(err); len(frames) != 0 {
				cfg = cfg.WithStackFrames(frames)
			}
		}
		return l.Output(fmt.Sprint(a[0]), cfg)
	}
	return l.Output(fmt.Sprint(a...), cfg.WithLevel(level).WithPositionSkip(2))
}
func (l *Logger) commonOutputf(level Level, format string, a ...any) error {
	a, cfg := spiltArgAndCfg(a, l.config)
	return l.Output(fmt.Sprintf(format, a...), cfg.WithLevel(level).WithPositionSkip(2))
}

func (l *Logger) Debug(a ...any) error { return l.commonOutput(LevelDebug, a...) }
func (l *Logger) Debugf(format string, a ...any) error {
	return l.commonOutputf(LevelDebug, format, a...)
}

func (l *Logger) Trace(a ...any) error { return l.commonOutput(LevelTrace, a...) }
func (l *Logger) Tracef(format string, a ...any) error {
	return l.commonOutputf(LevelTrace, format, a...)
}

func (l *Logger) Info(a ...any) error { return l.commonOutput(LevelInfo, a...) }
func (l *Logger) Infof(format string, a ...any) error {
	return l.commonOutputf(LevelInfo, format, a...)
}

func (l *Logger) Warn(a ...any) error { return l.commonOutput(LevelWarn, a...) }
func (l *Logger) Warnf(format string, a ...any) error {
	return l.commonOutputf(LevelWarn, format, a...)
}

func (l *Logger) Keyword(a ...any) error { return l.commonOutput(LevelKeyword, a...) }
func (l *Logger) Keywordf(format string, a ...any) error {
	return l.commonOutputf(LevelKeyword, format, a...)
}

func (l *Logger) Error(a ...any) error { return l.commonOutput(LevelError, a...) }
func (l *Logger) Errorf(format string, a ...any) error {
	return l.commonOutputf(LevelError, format, a...)
}

func (l *Logger) Panic(a ...any) error { return l.commonOutput(LevelPanic, a...) }
func (l *Logger) Panicf(format string, a ...any) error {
	return l.commonOutputf(LevelPanic, format, a...)
}
