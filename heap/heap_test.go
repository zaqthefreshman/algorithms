package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapify(t *testing.T) {
	a := []int{1, 3, 7, 2, 0, 9, 4}

	h := Heap(a)

	t.Log(a)
	t.Log(h)
	require.Equal(t, h[0], 9)
	require.NotEqual(t, h[0], a[0])
}
