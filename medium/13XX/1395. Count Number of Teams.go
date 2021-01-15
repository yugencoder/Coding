package t13XX

import "sort"

/*1395. Count Number of Teams

There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
A team is valid if:  (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).
Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).



Example 1:

Input: rating = [2,5,3,4,1]
Output: 3
Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1).
Example 2:

Input: rating = [2,1,3]
Output: 0
Explanation: We can't form any team given the conditions.
Example 3:

Input: rating = [1,2,3,4]
Output: 4


Constraints:

n == rating.length
1 <= n <= 200
1 <= rating[i] <= 10^5
*/

func update(bit []int, n int, value int) {
	for ; n < len(bit); n += n & -n {
		bit[n] += value
	}
}

func query(bit []int, n int) int {
	res := 0
	for ; n > 0; n -= n & -n {
		res += bit[n]
	}

	return res
}

func NumTeams(rating []int) int {
	res := 0
	n := len(rating)
	bit := make([]int, n+1)
	mapping := map[int]int{}

	lfSmNum, lfGtNum := make([]int, len(rating)), make([]int, len(rating))
	rtSmNum, rtGtNum := make([]int, len(rating)), make([]int, len(rating))

	sorted := make([]int, n)
	copy(sorted, rating)
	sort.Ints(sorted)

	for i, r := range sorted {
		mapping[r] = i + 1
		//note - i+1 usage
	}

	for i := 0; i < n; i++ {
		lfSmNum[i] = query(bit, mapping[rating[i]])
		lfGtNum[i] = query(bit, n) - query(bit, mapping[rating[i]])
		update(bit, mapping[rating[i]], 1)
	}

	bit = make([]int, n+1)

	for i := n - 1; i > 0; i-- {
		rtSmNum[i] = query(bit, mapping[rating[i]])
		rtGtNum[i] = query(bit, n) - query(bit, mapping[rating[i]])
		update(bit, mapping[rating[i]], 1)
		res += lfSmNum[i]*rtGtNum[i] + lfGtNum[i]*rtSmNum[i]
	}

	return res
}

// o(nlogn)
func NumTeams_On2(rating []int) int {
	res := 0
	leftLessArr := make([]int, len(rating))
	rightMoreArr := make([]int, len(rating))

	leftMoreArr := make([]int, len(rating))
	rightLessArr := make([]int, len(rating))

	for i := range rating {
		for j := range rating {
			if rating[j] < rating[i] {
				if j < i {
					leftLessArr[i]++
				} else {
					rightLessArr[i]++
				}
			}

			if rating[j] > rating[i] {
				if j < i {
					leftMoreArr[i]++
				} else {
					rightMoreArr[i]++
				}
			}
		}
	}

	for i := range rating {
		res += leftLessArr[i]*rightMoreArr[i] + leftMoreArr[i]*rightLessArr[i]
	}

	return res
}

func numTeams_On3(rating []int) int {
	res := 0
	for i := range rating {
		for j := 0; j < i; j++ {
			for k := j + 1; k < len(rating); k++ {
				if rating[j] > rating[i] && rating[i] > rating[k] {
					res++
				}

				if rating[j] < rating[i] && rating[i] < rating[k] {
					res++
				}
			}

		}

	}
	return res
}
