package problem

/*
Given an integer array nums, return the sum of divisors of the integers in that array that have exactly four divisors.

If there is no such integer in the array, return 0.



Example 1:

Input: nums = [21,4,7]
Output: 32
Explanation:
21 has 4 divisors: 1, 3, 7, 21
4 has 3 divisors: 1, 2, 4
7 has 2 divisors: 1, 7
The answer is the sum of divisors of 21 only.


Constraints:

1 <= nums.length <= 10^4
1 <= nums[i] <= 10^5
*/

func getAllPrimes(n int) ([]bool, []int) {
	primes := make([]bool, n)
	primeNos := []int{}

	for i := range primes {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		for j := i + i; j < n; j += i {
			primes[j] = false
		}
	}
	for i, p := range primes {
		if p && i > 1 {
			primeNos = append(primeNos, i)
		}
	}
	return primes, primeNos
}

func SumFourDivisors(nums []int) int {
	primeIdx, primes := getAllPrimes(100010)
	sum := 0

	for _, n := range nums {
		noDivisor := 2
		divisorSum := 1 + n

		for _, p := range primes {
			if p >= n || noDivisor > 4 {
				break
			}

			if n%p == 0 {
				otherDivisor := n / p

				if otherDivisor == p*p || primeIdx[otherDivisor] {
					divisorSum += p
					noDivisor++

					if otherDivisor == p*p {
						divisorSum += otherDivisor
						noDivisor++
					}

				} else {
					break
				}

			}

		}

		if noDivisor == 4 {
			sum += divisorSum
		}
	}

	return sum
}
