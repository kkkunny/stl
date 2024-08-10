package stllog

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// 是否是终端
func writerIsTerminal(w io.Writer) bool {
	file, ok := w.(*os.File)
	if !ok {
		return false
	}
	return isatty.IsTerminal(file.Fd()) || isatty.IsCygwinTerminal(file.Fd())
}
