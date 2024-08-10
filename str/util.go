package stlstr

import "strings"

// centerAlignString 将字符串进行居中对齐，使用空格填充
func centerAlignString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	// 计算左侧和右侧需要填充的空格数
	totalPadding := width - len(s)
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding
	return strings.Repeat(" ", leftPadding) + s + strings.Repeat(" ", rightPadding)
}

// CenterAlignStrings 对字符串切片进行居中对齐
func CenterAlignStrings(strings []string) []string {
	// 找到最长的字符串长度
	maxLen := 0
	for _, s := range strings {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	// 对每个字符串进行居中对齐
	alignedStrings := make([]string, len(strings))
	for i, s := range strings {
		alignedStrings[i] = centerAlignString(s, maxLen)
	}
	return alignedStrings
}
