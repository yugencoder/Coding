package sorting

func BubbleSort(items []int) {
	l := len(items)

	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

// can add a swap to break easily
