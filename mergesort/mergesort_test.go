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
	sorted, _ := MergeSort(input)
	assert.Equal(t, expected, sorted)
}

func TestMergeSortCount1(t *testing.T) {
	input := []int{2, 6, 1, 4, 5}
	expectedInversions := 4

	_, inversions := MergeSort(input)
	assert.Equal(t, expectedInversions, inversions)
}

func TestMergeSortCount2(t *testing.T) {
	input := []int{1, 20, 6, 4, 5}
	expectedInversions := 5

	_, inversions := MergeSort(input)
	assert.Equal(t, expectedInversions, inversions)
}
