package stllog

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := DefaultLogger(true)
	l := logger.NewGroup("111")
	l.Debug("hello")
	l.SDebugStack(0, 1)
	l.Info("hello")
	l.Warn("hello")
	l.Keyword("hello")
	l.Error("hello")
	l.Panicf("hello")
}
