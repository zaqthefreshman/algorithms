package mergesort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(rand.Perm(100001))
	}
}

func TestMergeSort(t *testing.T) {
	expected := make([]int, 0)
	for i := 0; i < 1000001; i++ {
		expected = append(expected, i)
	}
	input := make([]int, len(expected))
	copy(input, expected)
	for i := range input {
		j := rand.Intn(i + 1)
		input[i], input[j] = input[j], input[i]
	}

	assert.Equal(t, expected, MergeSort(input))
}
