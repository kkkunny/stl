package stlsync

import (
	"strings"
	"sync"
	"testing"
	"time"

	stltest "github.com/kkkunny/stl/test"
)

func TestReentrantLock_SimpleLock(t *testing.T) {
	locker := NewReentrantLock()
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

func TestReentrantLock_ReLock(t *testing.T) {
	locker := NewReentrantLock()

	locker.Lock()
	locker.Lock()
	locker.Unlock()
	locker.Unlock()

	stltest.AssertEq(t, 1, 1)
}

func TestReentrantRWLock_SimpleLock(t *testing.T) {
	locker := NewReentrantRWLock()
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

func TestReentrantRWLock_ReLock(t *testing.T) {
	locker := NewReentrantRWLock()

	locker.Lock()
	locker.Lock()
	locker.Unlock()
	locker.Unlock()

	stltest.AssertEq(t, 1, 1)
}

func TestReentrantRWLock_SimpleRLock(t *testing.T) {
	locker := NewReentrantRWLock()
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

func TestReentrantRWLock_ReRLock(t *testing.T) {
	locker := NewReentrantRWLock()

	locker.RLock()
	locker.RLock()
	locker.RUnlock()
	locker.RUnlock()

	stltest.AssertEq(t, 1, 1)
}

func TestReentrantRWLock_DowngradeLock(t *testing.T) {
	locker := NewReentrantRWLock()

	var s strings.Builder

	var wg sync.WaitGroup
	wg.Go(func() {
		locker.Lock()
		locker.RLock()
		defer locker.RUnlock()
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
