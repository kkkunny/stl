package stllog

import (
	"context"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlos "github.com/kkkunny/stl/os"
)

type config struct {
	ctx   context.Context
	level Level

	displayLevel      bool        // 是否输出日志级别
	displayTimeFormat string      // 日志打印时间格式
	displayPos        bool        // 是否输出日志打印位置
	displayPosFrame   stlos.Frame // 日志打印位置
	posSkip           uint
	displayColor      bool          // 是否对日志进行染色
	displayStack      bool          // 是否输出堆栈
	displayGroup      bool          // 是否输出日志组别
	frames            []stlos.Frame // 堆栈
}

func spiltArgAndCfg(a []any, defaultCfg config) ([]any, config) {
	as := make([]any, 0, len(a))
	cfgs := make([]config, 0, len(a))
	for _, aa := range a {
		switch cfg := aa.(type) {
		case config:
			cfgs = append(cfgs, cfg)
		default:
			as = append(as, aa)
		}
	}
	return as, stlslices.Last(cfgs, defaultCfg)
}

// WithLevel 设置日志等级
func (cfg config) WithLevel(level Level) config {
	cfg.level = level
	return cfg
}

// WithContext 设置上下文
func (cfg config) WithContext(ctx context.Context) config {
	cfg.ctx = ctx
	return cfg
}

// WithDisplayLevel 显示等级
func (cfg config) WithDisplayLevel() config {
	cfg.displayLevel = true
	return cfg
}

// WithNoDisplayLevel 禁止显示等级
func (cfg config) WithNoDisplayLevel() config {
	cfg.displayLevel = false
	return cfg
}

// WithDisplayTime 显示时间
func (cfg config) WithDisplayTime(format ...string) config {
	cfg.displayTimeFormat = stlslices.Last(format, time.DateTime)
	return cfg
}

// WithDisplayPosition 显示代码位置
func (cfg config) WithDisplayPosition(frame ...stlos.Frame) config {
	cfg.displayPos = true
	if len(frame) > 0 {
		cfg.displayPosFrame = frame[len(frame)-1]
	}
	return cfg
}

// WithNoDisplayPosition 禁止显示代码位置
func (cfg config) WithNoDisplayPosition() config {
	cfg.displayPos = false
	cfg.displayPosFrame = nil
	return cfg
}

// WithPositionSkip 代码位置跳过次数
func (cfg config) WithPositionSkip(skip uint) config {
	cfg.posSkip += skip
	return cfg
}

// WithDisplayColor 显示颜色
func (cfg config) WithDisplayColor() config {
	cfg.displayColor = true
	return cfg
}

// WithNoDisplayColor 禁止显示颜色
func (cfg config) WithNoDisplayColor() config {
	cfg.displayColor = false
	return cfg
}

// WithDisplayStack 显示堆栈信息
func (cfg config) WithDisplayStack() config {
	cfg.displayStack = true
	return cfg
}

// WithNoDisplayStack 禁止显示堆栈信息
func (cfg config) WithNoDisplayStack() config {
	cfg.displayStack = false
	return cfg
}

// WithStackFrames 堆栈信息
func (cfg config) WithStackFrames(frames []stlos.Frame) config {
	cfg.frames = frames
	return cfg
}

// WithDisplayGroup 显示组信息
func (cfg config) WithDisplayGroup() config {
	cfg.displayGroup = true
	return cfg
}

// WithNoDisplayGroup 禁止显示组信息
func (cfg config) WithNoDisplayGroup() config {
	cfg.displayGroup = false
	return cfg
}
