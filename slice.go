package gentypes

import "sort"

// SliceTransform is a function that transforms a slice of type T into a slice of type U.
func SliceTransform[T1 any, T2 any](slice []T1, f func(T1) T2) []T2 {
	result := make([]T2, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}

type sortInterface[T any] struct {
	slice    []T
	lessThan func(a, b T) bool
}

func (s *sortInterface[T]) Len() int {
	return len(s.slice)
}

func (s *sortInterface[T]) Less(i int, j int) bool {
	return s.lessThan(s.slice[i], s.slice[j])
}

// Swap swaps the elements with indexes i and j.
func (s *sortInterface[T]) Swap(i int, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func SliceSortWithComparison[T any](slice []T, comparison func(a, b T) bool) []T {
	sort.Sort(&sortInterface[T]{slice: slice, lessThan: comparison})
	return slice
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func SliceSort[T Ordered](slice []T) []T {
	return SliceSortWithComparison(slice, func(a, b T) bool {
		return a < b
	})
}
