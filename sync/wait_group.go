package sync

import (
	"context"
	"golang.org/x/sync/errgroup"
	"testing/iotest"
	"time"
)

// WaitGroup 等待组
type WaitGroup struct {
	ctx   context.Context
	group *errgroup.Group
}

// NewWaitGroup 新建等待组
func NewWaitGroup(ctx context.Context, limit ...uint) *WaitGroup {
	if ctx == nil {
		ctx = context.Background()
	}

	group, ctx := errgroup.WithContext(ctx)
	if len(limit) > 0 && limit[0] > 0 {
		group.SetLimit(int(limit[0]))
	}

	return &WaitGroup{
		ctx:   ctx,
		group: group,
	}
}

// Add 增加任务
func (self *WaitGroup) Add(t func(ctx context.Context) error) {
	self.group.Go(func() error {
		return t(self.ctx)
	})
}

// TryAdd 尝试增加任务
func (self *WaitGroup) TryAdd(t func(ctx context.Context) error) {
	self.group.TryGo(func() error {
		return t(self.ctx)
	})
}

// Wait 等待
func (self *WaitGroup) Wait() error {
	return self.group.Wait()
}

// WaitWithTimeOut 超时等待
func (self *WaitGroup) WaitWithTimeOut(d time.Duration) (err error) {
	timeout := time.NewTimer(d)
	endCh := make(chan error)

	go func() {
		waitErr := self.group.Wait()
		endCh <- waitErr
	}()

	select {
	case <-timeout.C:
		return iotest.ErrTimeout
	case err = <-endCh:
		return err
	}
}
