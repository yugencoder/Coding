package easy

import "fmt"

type TwoSum struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{
		nums: make([]int, 0),
	}

}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {

	if len(this.nums) == 0 {
		this.nums = []int{number}
		fmt.Println(this.nums)

		return
	}

	newNums := make([]int, len(this.nums)+1)
	for i, n := range this.nums {
		if number <= n {
			copy(newNums[:i], this.nums[:i])
			newNums[i] = number
			copy(newNums[i+1:], this.nums[i:])
			break

		} else if i == len(this.nums)-1 {
			newNums[i+1] = number
		}
	}

	this.nums = newNums

	fmt.Println(this.nums)

}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	if len(this.nums) == 0 {
		return false
	}

	i := 0
	j := len(this.nums) - 1
	for i < j && j >= 0 && i < len(this.nums) {
		sum := this.nums[i] + this.nums[j]
		if sum == value {
			return true
		}

		if sum < value {
			i++
		} else {
			j--
		}

	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
