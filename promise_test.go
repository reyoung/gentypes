package gentypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	promise := NewPromise(func() (int, error) {
		return 10, nil
	})
	promise.Apply()
	status := StatusFlatten(ChanHead(promise.Result()))
	require.Equal(t, 10, status.Value())
}
