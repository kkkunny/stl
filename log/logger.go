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

	stlerror "github.com/kkkunny/stl/error"
)

// LogLevel 日志等级
type LogLevel uint8

const (
	LogLevelDebug   LogLevel = iota // debug
	LogLevelInfo                    // info
	LogLevelWarn                    // warn
	LogLevelKeyword                 // keyword
	LogLevelError                   // error
)

var logLevelStringMap = [...]string{
	LogLevelDebug:   "  DEBUG  ",
	LogLevelInfo:    "   INFO  ",
	LogLevelWarn:    "   WARN  ",
	LogLevelKeyword: " KEYWORD ",
	LogLevelError:   "  ERROR  ",
}

var logLevelColorMap = [...]color.Color{
	LogLevelDebug:   color.Blue,
	LogLevelInfo:    color.Green,
	LogLevelWarn:    color.Yellow,
	LogLevelError:   color.Red,
	LogLevelKeyword: color.Magenta,
}

var logLevelStyleMap = [...]color.Style{
	LogLevelDebug:   color.New(color.OpBold, color.White, color.BgBlue),
	LogLevelInfo:    color.New(color.OpBold, color.White, color.BgGreen),
	LogLevelWarn:    color.New(color.OpBold, color.White, color.BgYellow),
	LogLevelError:   color.New(color.OpBold, color.White, color.BgRed),
	LogLevelKeyword: color.New(color.OpBold, color.White, color.BgMagenta),
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

// 输出
func (self *Logger) output(level LogLevel, pos string, msg string) error {
	var routesBuf strings.Builder
	for _, r := range self.routes {
		routesBuf.WriteString(" | ")
		routesBuf.WriteString(r)
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	var s string
	writer := self.writer.Writer()
	if writer == os.Stdout || writer == os.Stderr {
		suffix := fmt.Sprintf(
			"| %s | %s%s | %s",
			timeStr,
			pos,
			routesBuf.String(),
			msg,
		)
		suffix = logLevelColorMap[level].Text(suffix)
		s = logLevelStyleMap[level].Sprintf(logLevelStringMap[level]) + suffix
	} else {
		s = fmt.Sprintf(
			"%s| %s | %s%s | %s",
			logLevelStringMap[level],
			timeStr,
			pos,
			routesBuf.String(),
			msg,
		)
	}
	return self.writer.Output(0, s)
}

func (self *Logger) outputByStack(level LogLevel, skip uint, msg string) error {
	_, file, line, _ := runtime.Caller(int(skip + 1))
	return self.output(level, fmt.Sprintf("%s:%d", file, line), msg)
}

// 打印
func (self *Logger) print(level LogLevel, skip uint, a any) error {
	if self.level > level {
		return nil
	}
	return self.outputByStack(level, skip+1, fmt.Sprintf("%v", a))
}

// 格式化打印
func (self *Logger) printf(level LogLevel, skip uint, f string, a ...any) error {
	return self.print(level, skip+1, fmt.Sprintf(f, a...))
}

// 打印异常
func (self *Logger) printError(level LogLevel, skip uint, err error) error {
	var logerr stlerror.Error
	if errors.As(err, &logerr) {
		return self.printLogError(level, logerr)
	} else {
		return self.print(level, skip+1, err.Error())
	}
}

// 打印带栈异常
func (self *Logger) printLogError(level LogLevel, err stlerror.Error) error {
	if self.level > level {
		return nil
	}

	stacks := err.Stacks()

	var stackBuffer strings.Builder
	for i, s := range stacks {
		stackBuffer.WriteString(fmt.Sprintf("\t%s:%d", s.File, s.Line))
		if i < len(stacks)-1 {
			stackBuffer.WriteByte('\n')
		}
	}

	stack := stacks[len(stacks)-1]
	return self.output(level, fmt.Sprintf("%s:%d", stack.File, stack.Line), fmt.Sprintf("%s\n%s", err.Error(), stackBuffer.String()))
}

// Debug 输出Debug信息
func (self *Logger) Debug(skip uint, a any) error {
	return self.print(LogLevelDebug, skip+1, a)
}

// Debugf 输出Debugf格式化信息
func (self *Logger) Debugf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelDebug, skip+1, f, a...)
}

// DebugError 输出Debug异常信息
func (self *Logger) DebugError(skip uint, err error) error {
	return self.printError(LogLevelDebug, skip+1, err)
}

// Info 输出Info信息
func (self *Logger) Info(skip uint, a any) error {
	return self.print(LogLevelInfo, skip+1, a)
}

// Infof 输出Info格式化信息
func (self *Logger) Infof(skip uint, f string, a ...any) error {
	return self.printf(LogLevelInfo, skip+1, f, a...)
}

// InfoError 输出Info异常信息
func (self *Logger) InfoError(skip uint, err error) error {
	return self.printError(LogLevelInfo, skip+1, err)
}

// Warn 输出Warn信息
func (self *Logger) Warn(skip uint, a any) error {
	return self.print(LogLevelWarn, skip+1, a)
}

// Warnf 输出Warn格式化信息
func (self *Logger) Warnf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelWarn, skip+1, f, a...)
}

// WarnError 输出Warn异常信息
func (self *Logger) WarnError(skip uint, err error) error {
	return self.printError(LogLevelWarn, skip+1, err)
}

// Error 输出Error信息
func (self *Logger) Error(skip uint, a any) error {
	return self.print(LogLevelError, skip+1, a)
}

// Errorf 输出Error格式化信息
func (self *Logger) Errorf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelError, skip+1, f, a...)
}

// ErrorError 输出Error异常信息
func (self *Logger) ErrorError(skip uint, err error) error {
	return self.printError(LogLevelError, skip+1, err)
}

// Keyword 输出Keyword信息
func (self *Logger) Keyword(skip uint, a any) error {
	return self.print(LogLevelKeyword, skip+1, a)
}

// Keywordf 输出Keyword格式化信息
func (self *Logger) Keywordf(skip uint, f string, a ...any) error {
	return self.printf(LogLevelKeyword, skip+1, f, a...)
}

// KeywordError 输出Keyword异常信息
func (self *Logger) KeywordError(skip uint, err error) error {
	return self.printError(LogLevelKeyword, skip+1, err)
}
