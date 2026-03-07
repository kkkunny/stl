package stlsync

import (
	"strings"
	"sync"
	"testing"
	"time"

	stltest "github.com/kkkunny/stl/test"
)

func BenchmarkMutex(b *testing.B) {
	m := sync.Mutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			m.Unlock()
		}
	})
}

func BenchmarkSpinLock(b *testing.B) {
	spin := NewSpinLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			spin.Unlock()
		}
	})
}

func TestSpinLock_Lock(t *testing.T) {
	locker := NewSpinLock()
	var s strings.Builder

	var wg sync.WaitGroup
	wg.Go(func() {
		locker.Lock()
		defer locker.Unlock()
		time.Sleep(time.Millisecond * 2)
		s.WriteString("1")
	})
	wg.Go(func() {
		time.Sleep(time.Millisecond)
		locker.Lock()
		defer locker.Unlock()
		s.WriteString("2")
	})
	wg.Wait()

	stltest.AssertEq(t, s.String(), "12")
}

func TestSpinRWLock_Lock(t *testing.T) {
	locker := NewSpinRWLock()
	var s strings.Builder

	var wg sync.WaitGroup
	wg.Go(func() {
		locker.Lock()
		defer locker.Unlock()
		time.Sleep(time.Millisecond * 2)
		s.WriteString("1")
	})
	wg.Go(func() {
		time.Sleep(time.Millisecond)
		locker.Lock()
		defer locker.Unlock()
		s.WriteString("2")
	})
	wg.Wait()

	stltest.AssertEq(t, s.String(), "12")
}

func TestSpinRWLock_RLock(t *testing.T) {
	locker := NewSpinRWLock()
	var s strings.Builder

	var wg sync.WaitGroup
	wg.Go(func() {
		locker.RLock()
		defer locker.RUnlock()
		time.Sleep(time.Millisecond * 2)
		s.WriteString("1")
	})
	wg.Go(func() {
		time.Sleep(time.Millisecond)
		locker.RLock()
		defer locker.RUnlock()
		s.WriteString("2")
	})
	wg.Wait()

	stltest.AssertEq(t, s.String(), "21")
}
