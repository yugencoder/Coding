package _8XX

/*
You are given an array nums that consists of non-negative integers. Let us define rev(x) as the reverse of the non-negative integer x. For example, rev(123) = 321, and rev(120) = 21. A pair of indices (i, j) is nice if it satisfies all of the following conditions:

0 <= i < j < nums.length
nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
Return the number of nice pairs of indices. Since that number can be too large, return it modulo 109 + 7.

*/

func countNicePairs(nums []int) int {
	var res int
	mod := int(1e9 + 7)
	aMap := map[int]int{}

	for i, n := range nums {
		aMap[getRev(n) - nums[i]] = (aMap[getRev(n) - nums[i]] + 1) % mod
	}

	for _, v := range aMap {
		res = (res%mod + (v*(v-1)/2)%mod) % mod
	}

	return res
}

func getRev(n int) int {
	rev := 0
	for ; n > 0; {
		rev = rev*10 + n%10
		n = n / 10
	}

	return rev
}
