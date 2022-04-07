package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// 日志等级
const (
	LEVEL_DEBUG int8 = iota // debug
	LEVEL_INFO              // info
	LEVEL_WARN              // warn
	LEVEL_ERROR             // error
	LEVEL_PANIC             // panic
)

// 日志管理器
type Logger struct {
	level  int8         // 等级
	mutex  sync.RWMutex // 锁
	writer io.Writer
}

var std = New(os.Stdout, LEVEL_DEBUG)

// 默认日志管理器
func Default() *Logger {
	return std
}

// 新建日志管理器
func New(writer io.Writer, level int8) *Logger {
	return &Logger{
		level:  level,
		writer: writer,
	}
}

// 设置等级
func (self *Logger) SetLevel(level int8) {
	if self.Level() == level {
		return
	}
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.level = level
}

// 获取等级
func (self *Logger) Level() int8 {
	self.mutex.RLock()
	defer self.mutex.RUnlock()
	return self.level
}

// 输出
func (self *Logger) output(prefix string, msg string, skip int, fcolor, bcolor int) (err error) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	_, file, line, _ := runtime.Caller(skip)
	if self.writer == os.Stdout || self.writer == os.Stderr {
		_, err = fmt.Fprintf(self.writer, "\033[1;%d;%dm %s \033[0m\033[1;%dm| %s | %s:%d | %s\033[0m\n", fcolor, bcolor, prefix, fcolor, timeStr, file, line, msg)
	} else {
		_, err = fmt.Fprintf(self.writer, " %s | %s | %s:%d | %s\n", prefix, timeStr, file, line, msg)
	}
	return err
}

// 输出debug
func (self *Logger) Debug(a ...any) error {
	if self.Level() > LEVEL_DEBUG {
		return nil
	}
	return self.output("DEBUG", fmt.Sprint(a...), 2, 36, 46)
}

// 格式化输出debug
func (self *Logger) Debugf(format string, a ...any) error {
	if self.Level() > LEVEL_DEBUG {
		return nil
	}
	return self.output("DEBUG", fmt.Sprintf(format, a...), 2, 36, 46)
}

// 输出info
func (self *Logger) Info(a ...any) error {
	if self.Level() > LEVEL_INFO {
		return nil
	}
	return self.output("INFO ", fmt.Sprint(a...), 2, 32, 42)
}

// 格式化输出info
func (self *Logger) Infof(format string, a ...any) error {
	if self.Level() > LEVEL_INFO {
		return nil
	}
	return self.output("INFO ", fmt.Sprintf(format, a...), 2, 32, 42)
}

// 输出warn
func (self *Logger) Warn(a ...any) error {
	if self.Level() > LEVEL_WARN {
		return nil
	}
	return self.output("WARN ", fmt.Sprint(a...), 2, 33, 43)
}

// 格式化输出warn
func (self *Logger) Warnf(format string, a ...any) error {
	if self.Level() > LEVEL_WARN {
		return nil
	}
	return self.output("WARN ", fmt.Sprintf(format, a...), 2, 33, 43)
}

// 输出error
func (self *Logger) Error(a ...any) error {
	if self.Level() > LEVEL_ERROR {
		return nil
	}
	return self.output("ERROR", fmt.Sprint(a...), 2, 31, 41)
}

// 格式化输出error
func (self *Logger) Errorf(format string, a ...any) error {
	if self.Level() > LEVEL_ERROR {
		return nil
	}
	return self.output("ERROR", fmt.Sprintf(format, a...), 2, 31, 41)
}

// 输出panic
func (self *Logger) Panic(a ...any) {
	if self.Level() > LEVEL_PANIC {
		return
	}
	s := fmt.Sprint(a...)
	_ = self.output("PANIC", s, 2, 35, 45)
	panic(s)
}

// 格式化输出panic
func (self *Logger) Panicf(format string, a ...any) {
	if self.Level() > LEVEL_PANIC {
		return
	}
	s := fmt.Sprintf(format, a...)
	_ = self.output("PANIC", s, 2, 35, 45)
	panic(s)
}
