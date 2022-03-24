package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	m := map[string]string{
		"a": "b",
		"c": "d",
	}
	require.Equal(t, []string{"a", "c"}, SliceSort(MapKeys(m)))
}
