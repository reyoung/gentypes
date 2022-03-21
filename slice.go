package gentypes

// SliceTransform is a function that transforms a slice of type T into a slice of type U.
func SliceTransform[T1 any, T2 any](slice []T1, f func(T1) T2) []T2 {
	result := make([]T2, 0, len(slice))
	for _, v := range slice {
		result = append(result, f(v))
	}
	return result
}
