package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require.Equal(t, StatusTransform(NewStatus(1, nil), func(t int) (int, error) {
		return t * 2, nil
	}).Value(), 2)
}
