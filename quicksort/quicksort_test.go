package quicksort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(rand.Perm(100001))
	}
}

func TestQuickSort(t *testing.T) {
	expected := make([]int, 0)
	for i := 0; i < 100001; i++ {
		expected = append(expected, i)
	}
	input := make([]int, len(expected))
	copy(input, expected)
	for i := range input {
		j := rand.Intn(i + 1)
		input[i], input[j] = input[j], input[i]
	}
	sorted := QuickSort(input)
	assert.Equal(t, expected, sorted)
}
