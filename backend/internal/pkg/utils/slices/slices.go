package slices

import (
	"cmp"
	"slices"
)

// DelDupls deletes duplicates from the given slice and returns updated slice.
func DelDupls[S ~[]E, E cmp.Ordered](s S) S {
	slices.Sort(s)
	return slices.Compact(s)
}

// DelDuplsFunc deletes duplicates from slice based on a key function.
func DelDuplsFunc[T any, K cmp.Ordered](slice []T, keyFunc func(T) K) []T {
	seen := make(map[K]bool)
	res := make([]T, 0, len(slice))

	var key K
	for _, item := range slice {
		key = keyFunc(item)
		if !seen[key] {
			seen[key] = true
			res = append(res, item)
		}
	}
	return res
}
