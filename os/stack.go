package stlos

import "runtime"

func GetCallStacks(depth uint, skip ...uint) []runtime.Frame {
	var skipVal uint
	if len(skip) > 0 {
		skipVal = skip[0]
	}

	var reverseStacks []runtime.Frame
	pcs := make([]uintptr, depth)

	n := runtime.Callers(int(skipVal)+1, pcs)
	frames := runtime.CallersFrames(pcs[:n-1])
	for frame, exist := frames.Next(); exist; frame, exist = frames.Next() {
		if !exist {
			break
		}
		reverseStacks = append(reverseStacks, frame)
	}

	stacks := make([]runtime.Frame, len(reverseStacks))
	for i, s := range reverseStacks {
		stacks[len(reverseStacks)-i-1] = s
	}
	return stacks
}
