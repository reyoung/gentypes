package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptional(t *testing.T) {
	opt := Some(10)
	require.False(t, opt.IsNone())
	require.True(t, opt.IsSome())
	require.Equal(t, 10, opt.Value())
	require.Equal(t, 10, opt.OrElse(20))
	require.Equal(t, 20, OptionalTransform(opt, func(x int) int { return x * 2 }).Value())
}
