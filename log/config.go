package stllog

import (
	"context"
	"runtime"
	"time"

	stlslices "github.com/kkkunny/stl/container/slices"
)

type config struct {
	ctx   context.Context
	level Level

	displayLevel      bool
	displayTimeFormat string
	displayPos        bool
	posSkip           uint
	displayColor      bool
	displayStack      bool
	displayGroup      bool
	stacks            []runtime.Frame
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
func (cfg config) WithDisplayPosition() config {
	cfg.displayPos = true
	return cfg
}

// WithNoDisplayPosition 禁止显示代码位置
func (cfg config) WithNoDisplayPosition() config {
	cfg.displayPos = false
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

// WithStacks 堆栈信息
func (cfg config) WithStacks(stacks []runtime.Frame) config {
	cfg.stacks = stacks
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
