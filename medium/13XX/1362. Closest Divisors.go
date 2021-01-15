package t13XX

import "math"

func ClosestDivisors(num int) []int {
	r := squareRoot(num + 2)

	for i := r; i >= 0; i++ {
		if (num+2)%i == 0 {
			return []int{i, (num + 2) / i}
		} else if (num+1)%i == 0 {
			return []int{i, (num + 1) / i}
		}
	}
	return nil
}


func squareRoot(num int) int {
	return int(math.Sqrt(float64(num)))
}
