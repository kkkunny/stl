package stlerr

import (
	"errors"
	"fmt"
	"testing"

	pkgErrors "github.com/pkg/errors"
	stltest "github.com/kkkunny/stl/test"
)

func TestGetErrorStackFrames_WithStack(t *testing.T) {
	err := Errorf("err")
	stltest.AssertNotEq(t, len(GetErrorStackFrames(err)), 0)
}

func TestGetErrorStackFrames_PkgErrors(t *testing.T) {
	err := pkgErrors.WithStack(pkgErrors.New("err"))
	stltest.AssertNotEq(t, len(GetErrorStackFrames(err)), 0)
}

func TestGetErrorStackFrames_Wrapped(t *testing.T) {
	err := fmt.Errorf("wrap: %w", pkgErrors.WithStack(pkgErrors.New("err")))
	stltest.AssertNotEq(t, len(GetErrorStackFrames(err)), 0)
}

func TestGetErrorStackFrames_Nil(t *testing.T) {
	stltest.AssertEq(t, GetErrorStackFrames(nil), nil)
	stltest.AssertEq(t, GetErrorStackFrames(errors.New("err")), nil)
}