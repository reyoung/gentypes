package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceTransform(t *testing.T) {
	result := SliceTransform([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	require.Equal(t, []int{2, 4, 6}, result)

	require.Equal(t, []string{"a", "b", "c"}, SliceSort([]string{"b", "c", "a"}))
}
