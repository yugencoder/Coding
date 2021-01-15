package t13XX

type bit123 struct {
	arr []int
}

func newBit123(l int) bit123 {
	return bit123{
		make([]int, l),
	}
}

func (b *bit123) update_123(idx, val int) {
	for i := idx; i < len(b.arr); i = i + i&(^i+1) {
		b.arr[i] += val
	}
}

func (b *bit123) getSum_123(idx int) int {
	sum := 0
	for i := idx; i > 0; i = i - i&(^i+1) {
		sum += b.arr[i]
	}
	return sum
}

func NumTimesAllBlue(light []int) int {
	res := 0
	bit := newBit123(len(light) + 1)

	for i := 1; i <= len(light); i++ {
		bit.update_123(i, 0)
	}

	for _, l := range light {
		bit.update_123(l, 1)
		actualSum := bit.getSum_123(l)
		totalSum := bit.getSum_123(len(light))
		if actualSum == l {
			if (totalSum == l) || totalSum == len(light) {
				res++
			} else {
				for i := l + 1; i < len(light); i++ {
					newSum := bit.getSum_123(i)
					if newSum != i {
						break
					}
					if totalSum == i {
						res++
						break
					}
				}
			}
		}
	}

	return res
}

// better solution
func numTimesAllBlue(light []int) int {
	ans, maxNumTillNow := 0, -1
	for i := 0; i < len(light); i++ {
		maxNumTillNow = max(light[i], maxNumTillNow)
		if maxNumTillNow == i+1 {
			ans += 1
		}
	}
	return ans
}
