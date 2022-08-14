package sync

import (
	"context"
	"github.com/kkkunny/stl/util"
	"testing/iotest"
	"time"
)

// GoroutinePool 协程池
type GoroutinePool struct {
	ctx    context.Context
	task   func(ctx context.Context) error
	cannel func()
	waitCh chan struct{}
	err    error
}

// NewGoroutinePool 新建协程池
func NewGoroutinePool(ctx context.Context, cap uint, task func(ctx context.Context) error) *GoroutinePool {
	if cap == 0 {
		panic("the capacity of pool must be greater than zero")
	}

	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cannel := context.WithCancel(ctx)

	ch := make(chan struct{}, cap)
	for i := uint(0); i < cap; i++ {
		ch <- struct{}{}
	}

	return &GoroutinePool{
		ctx:    ctx,
		task:   task,
		cannel: cannel,
		waitCh: ch,
	}
}

// NewGoroutinePoolWithTimeout 新建带超时协程池
func NewGoroutinePoolWithTimeout(ctx context.Context, cap uint, task func(ctx context.Context) error, duration time.Duration) *GoroutinePool {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, _ = context.WithTimeout(ctx, duration)
	return NewGoroutinePool(ctx, cap, task)
}

// Run 运行
func (self *GoroutinePool) Run() error {
loop:
	for {
		select {
		case <-self.ctx.Done():
			close(self.waitCh)
			if self.err == nil {
				self.err = iotest.ErrTimeout
			}
			break loop
		case _ = <-self.waitCh:
			go func() {
				taskErr := self.task(self.ctx)
				if taskErr != nil {
					self.err = taskErr
					self.cannel()
				} else if !util.IsChanClosed(self.waitCh) {
					self.waitCh <- struct{}{}
				}
			}()
		}
	}
	return self.err
}
