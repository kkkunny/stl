//go:build go1.21

package slices

import "slices"

func Clone[S interface{ ~[]E }, E any](s S) S {
	return slices.Clone[S, E](s)
}

func SortFunc[S interface{ ~[]E }, E any](x S, cmp func(a E, b E) int) {
	slices.SortFunc[S, E](x, cmp)
}

func Reverse[S interface{ ~[]E }, E any](s S) {
	slices.Reverse[S, E](s)
}
