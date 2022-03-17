package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	promise := NewPromise(func() int {
		return 10
	})
	promise.Apply()
	status := ChanHead(promise.Result())
	require.Equal(t, 10, status.Value())
}
