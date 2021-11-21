package _8XX

//You are given a sorted array nums of n non-negative integers and an integer maximumBit. You want to perform the following query n times:
//
//Find a non-negative integer k < 2maximumBit such that nums[0] XOR nums[1] XOR ... XOR nums[nums.length-1] XOR k is maximized. k is the answer to the ith query.
//Remove the last element from the current array nums.
//Return an array answer, where answer[i] is the answer to the ith query.

func getMaximumXor(nums []int, maximumBit int) []int {
	res := make([]int, len(nums))
	xor := 1<<maximumBit - 1
	for i, n := range nums {
		xor = xor ^ n
		res[len(nums)-i-1] = xor
	}

	return res
}
