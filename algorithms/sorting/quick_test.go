package sorting

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{4, 5, 9, 2, 6, 8, 2, 1, 7, 3}

	QuickSort(arr)
	fmt.Println(arr)
	//assert.Equal(t, arr, []int{4, 5, 9, 2, 6, 8, 2, 1, 7, 3})

}
