package stllog

import (
	"strings"

	"github.com/gookit/color"

	stlslices "github.com/kkkunny/stl/container/slices"
	stlstr "github.com/kkkunny/stl/str"
	stlbasic "github.com/kkkunny/stl/value"
)

type Level uint8

var (
	LevelDebug   = newLevel("DEBUG", color.New(color.OpBold, color.FgCyan, color.BgLightBlue))
	LevelTrace   = newLevel("TRACE", color.New(color.OpBold, color.FgBlue, color.BgCyan))
	LevelInfo    = newLevel("INFO", color.New(color.OpBold, color.FgWhite, color.BgGreen))
	LevelWarn    = newLevel("WARN", color.New(color.OpBold, color.FgBlue, color.BgLightYellow))
	LevelKeyword = newLevel("KEYWORD", color.New(color.OpReset, color.FgLightCyan, color.BgLightMagenta))
	LevelError   = newLevel("ERROR", color.New(color.OpBold, color.FgRed, color.BgMagenta))
	LevelPanic   = newLevel("PANIC", color.New(color.OpBold, color.FgMagenta, color.BgRed))
)

var maxLevel Level
var levelStringSilce []string
var levelStyleSilce []color.Style

func newLevel(s string, color color.Style) Level {
	levelStringSilce = append(levelStringSilce, s)
	levelStyleSilce = append(levelStyleSilce, color)
	levelStringSilce = stlstr.CenterAlignStrings(levelStringSilce)
	maxLevel++
	return maxLevel - 1
}

func (lvl Level) String() string {
	return strings.TrimSpace(lvl.AlignString())
}

func (lvl Level) AlignString() string {
	return levelStringSilce[lvl]
}

func (lvl Level) Style() color.Style {
	return levelStyleSilce[lvl]
}

func (lvl Level) MsgColor() color.Color {
	return stlbasic.IgnoreWith(stlslices.FindLast(lvl.Style(), func(_ int, clr color.Color) bool {
		return clr.IsBg()
	})).ToFg()
}
