package _7XX

import (
	"math"
	"sort"
)

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	toppingCombinations := []int{}
	toppingCosts = append(toppingCosts, toppingCosts...)
	tMap := map[int]bool{}
	for i := 0; i < int(math.Pow(2, float64(len(toppingCosts)))); i++ {
		j := 0
		sum := 0
		for m := i; j < len(toppingCosts); m >>= 1 {
			k := 1 & m
			if k == 1 {
				sum += toppingCosts[j]
			}
			j++
		}
		if _, ok := tMap[sum]; !ok {
			toppingCombinations = append(toppingCombinations, sum)
		}

		tMap[sum] = true
	}

	totalCost := []int{}

	for _, b := range baseCosts {
		for _, t := range toppingCombinations {
			totalCost = append(totalCost, b+t)
		}
	}

	sort.Ints(totalCost)

	minDiff := math.MaxInt32
	val := 0
	prevDiff := minDiff
	for _, t := range totalCost {
		currDiff := int(math.Abs(float64(t - target)))

		if currDiff > prevDiff {
			return val
		}

		if currDiff < minDiff {
			minDiff = currDiff
			val = t
		}

		prevDiff = currDiff
	}

	return val
}
