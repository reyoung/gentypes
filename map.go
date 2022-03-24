package gentypes

func MapKeys[T1 comparable, T2 any](m map[T1]T2) []T1 {
	if len(m) == 0 {
		return nil
	}
	keys := make([]T1, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
