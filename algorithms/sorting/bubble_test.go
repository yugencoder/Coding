package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{6, 5, 3, 1, 8, 7, 2, 4, 2}
	BubbleSort(arr)
	assert.Equal(t, []int{1, 2, 2, 3, 4, 5, 6, 7, 8}, arr)
}
