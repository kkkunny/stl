package stllog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"

	stlbasic "github.com/kkkunny/stl/basic"
	stlerror "github.com/kkkunny/stl/error"
	stlos "github.com/kkkunny/stl/os"
)

// LogLevel 日志等级
type LogLevel uint8

const (
	LogLevelDebug   LogLevel = iota // debug
	LogLevelInfo                    // info
	LogLevelWarn                    // warn
	LogLevelKeyword                 // keyword
	LogLevelError                   // error
	LogLevelPanic                   // panic
)

var logLevelStringMap = [...]string{
	LogLevelDebug:   "  DEBUG  ",
	LogLevelInfo:    "   INFO  ",
	LogLevelWarn:    "   WARN  ",
	LogLevelKeyword: " KEYWORD ",
	LogLevelError:   "  ERROR  ",
	LogLevelPanic:   "  PANIC  ",
}

var logLevelColorMap = [...]color.Color{
	LogLevelDebug:   color.Blue,
	LogLevelInfo:    color.Green,
	LogLevelWarn:    color.Yellow,
	LogLevelKeyword: color.Cyan,
	LogLevelError:   color.Magenta,
	LogLevelPanic:   color.Red,
}

var logLevelStyleMap = [...]color.Style{
	LogLevelDebug:   color.New(color.OpBold, color.White, color.BgBlue),
	LogLevelInfo:    color.New(color.OpBold, color.White, color.BgGreen),
	LogLevelWarn:    color.New(color.OpBold, color.White, color.BgYellow),
	LogLevelKeyword: color.New(color.OpBold, color.White, color.BgCyan),
	LogLevelError:   color.New(color.OpBold, color.White, color.BgMagenta),
	LogLevelPanic:   color.New(color.OpBold, color.White, color.BgRed),
}

// Logger 日志管理器
type Logger struct {
	level  LogLevel
	routes []string
	writer *log.Logger
}

// DefaultLogger 默认日志管理器
func DefaultLogger(debug bool) *Logger {
	if debug {
		return NewLogger(LogLevelDebug, os.Stdout)
	}
	return NewLogger(LogLevelInfo, os.Stdout)
}

// NewLogger 新建日志管理器
func NewLogger(level LogLevel, writer io.Writer) *Logger {
	return &Logger{
		level:  level,
		writer: log.New(writer, "", 0),
	}
}

func (self *Logger) NewGroup(name string) *Logger {
	routes := make([]string, len(self.routes)+1)
	copy(routes, self.routes)
	routes[len(routes)-1] = name
	return &Logger{
		level:  self.level,
		routes: routes,
		writer: self.writer,
	}
}

func (self *Logger) Group() string {
	if len(self.routes) == 0 {
		return ""
	}
	return self.routes[len(self.routes)-1]
}

// 输出
func (self *Logger) output(level LogLevel, pos string, msg string) error {
	var routesBuf strings.Builder
	for _, r := range self.routes {
		routesBuf.WriteString(" | ")
		routesBuf.WriteString(r)
	}

	timeStr := time.Now().Format(time.DateTime)
	var s string
	writer := self.writer.Writer()
	if writer == os.Stdout || writer == os.Stderr {
		suffix := fmt.Sprintf(
			"| %s%s | %s | %s",
			timeStr,
			routesBuf.String(),
			pos,
			msg,
		)
		suffix = logLevelColorMap[level].Text(suffix)
		s = logLevelStyleMap[level].Sprintf(logLevelStringMap[level]) + suffix
	} else {
		s = fmt.Sprintf(
			"%s | %s%s | %s | %s",
			logLevelStringMap[level],
			timeStr,
			routesBuf.String(),
			pos,
			msg,
		)
	}
	return self.writer.Output(0, s)
}

func (self *Logger) outputWithPos(level LogLevel, skip uint, msg string) error {
	_, file, line, _ := runtime.Caller(int(skip + 1))
	return self.output(level, fmt.Sprintf("%s:%d", file, line), msg)
}

// 打印
func (self *Logger) print(level LogLevel, skip uint, a any) error {
	if self.level > level {
		return nil
	}
	return self.outputWithPos(level, skip+1, fmt.Sprintf("%v", a))
}

// 带栈打印
func (self *Logger) printWithStack(level LogLevel, skip uint, a any) error {
	if self.level > level {
		return nil
	}

	var stacks []runtime.Frame
	var stlerr stlerror.Error
	if err, ok := a.(error); ok && errors.As(err, &stlerr) {
		stacks = stlerr.Stacks()
	} else {
		stacks = stlos.GetCallStacks(20, skip+2)
	}

	var stackBuffer strings.Builder
	for i, s := range stacks {
		stackBuffer.WriteString(fmt.Sprintf("\t%s:%d", s.File, s.Line))
		if i < len(stacks)-1 {
			stackBuffer.WriteByte('\n')
		}
	}
	return self.outputWithPos(level, skip+1, fmt.Sprintf("%v\n%s", a, stackBuffer.String()))
}

// 格式化打印
func (self *Logger) printf(level LogLevel, skip uint, f string, a ...any) error {
	return self.print(level, skip+1, fmt.Sprintf(f, a...))
}

// SDebug 输出Debug信息
func (self *Logger) SDebug(skip uint, a any) error {
	return self.print(LogLevelDebug, skip+1, a)
}

// SDebugf 输出Debugf格式化信息
func (self *Logger) SDebugf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelDebug, skip+1, f, a...)
}

// SDebugStack 输出Debug异常信息
func (self *Logger) SDebugStack(skip uint, a any) error {
	return self.printWithStack(LogLevelDebug, skip+1, a)
}

// SInfo 输出Info信息
func (self *Logger) SInfo(skip uint, a any) error {
	return self.print(LogLevelInfo, skip+1, a)
}

// SInfof 输出Info格式化信息
func (self *Logger) SInfof(skip uint, f string, a ...any) error {
	return self.printf(LogLevelInfo, skip+1, f, a...)
}

// SInfoStack 输出Info异常信息
func (self *Logger) SInfoStack(skip uint, a any) error {
	return self.printWithStack(LogLevelInfo, skip+1, a)
}

// SWarn 输出Warn信息
func (self *Logger) SWarn(skip uint, a any) error {
	return self.print(LogLevelWarn, skip+1, a)
}

// SWarnf 输出Warn格式化信息
func (self *Logger) SWarnf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelWarn, skip+1, f, a...)
}

// SWarnStack 输出Warn异常信息
func (self *Logger) SWarnStack(skip uint, a any) error {
	return self.printWithStack(LogLevelWarn, skip+1, a)
}

// SKeyword 输出Keyword信息
func (self *Logger) SKeyword(skip uint, a any) error {
	return self.print(LogLevelKeyword, skip+1, a)
}

// SKeywordf 输出Keyword格式化信息
func (self *Logger) SKeywordf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelKeyword, skip+1, f, a...)
}

// SKeywordStack 输出Keyword异常信息
func (self *Logger) SKeywordStack(skip uint, a any) error {
	return self.printWithStack(LogLevelKeyword, skip+1, a)
}

// SError 输出Error信息
func (self *Logger) SError(skip uint, a any) error {
	return self.print(LogLevelError, skip+1, a)
}

// SErrorf 输出Error格式化信息
func (self *Logger) SErrorf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelError, skip+1, f, a...)
}

// SErrorStack 输出Error异常信息
func (self *Logger) SErrorStack(skip uint, a any) error {
	return self.printWithStack(LogLevelError, skip+1, a)
}

// SPanic 输出Panic信息
func (self *Logger) SPanic(skip uint, a any) error {
	return self.print(LogLevelPanic, skip+1, a)
}

// SPanicf 输出Panic格式化信息
func (self *Logger) SPanicf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelPanic, skip+1, f, a...)
}

// SPanicStack 输出Panic异常信息
func (self *Logger) SPanicStack(skip uint, a any) error {
	return self.printWithStack(LogLevelPanic, skip+1, a)
}

func (self *Logger) Debug(a any)                 { stlbasic.Ignore(self.SDebug(1, a)) }
func (self *Logger) Debugf(f string, a ...any)   { stlbasic.Ignore(self.SDebugf(1, f, a...)) }
func (self *Logger) Info(a any)                  { stlbasic.Ignore(self.SInfo(1, a)) }
func (self *Logger) Infof(f string, a ...any)    { stlbasic.Ignore(self.SInfof(1, f, a...)) }
func (self *Logger) Warn(a any)                  { stlbasic.Ignore(self.SWarn(1, a)) }
func (self *Logger) Warnf(f string, a ...any)    { stlbasic.Ignore(self.SWarnf(1, f, a...)) }
func (self *Logger) Keyword(a any)               { stlbasic.Ignore(self.SKeyword(1, a)) }
func (self *Logger) Keywordf(f string, a ...any) { stlbasic.Ignore(self.SKeywordf(1, f, a...)) }
func (self *Logger) Error(a any)                 { stlbasic.Ignore(self.SErrorStack(1, a)) }
func (self *Logger) Errorf(f string, a ...any)   { stlbasic.Ignore(self.SErrorf(1, f, a...)) }
func (self *Logger) Panic(a any)                 { stlbasic.Ignore(self.SPanicStack(1, a)) }
func (self *Logger) Panicf(f string, a ...any)   { stlbasic.Ignore(self.SPanicf(1, f, a...)) }
