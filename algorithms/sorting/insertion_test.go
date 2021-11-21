package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 2}
	insertionSort(arr)
	assert.Equal(t, []int{1, 2, 2, 3, 4, 5, 6, 7, 8}, arr)
}
