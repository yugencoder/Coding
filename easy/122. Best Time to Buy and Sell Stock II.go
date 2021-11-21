package easy

func MaxProfit(prices []int) int {
	profits := make([]int, len(prices))
	profits[0] = 0

	for i := 0; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			currProfit := 0

			if prices[i] > prices[j] {
				currProfit = prices[i] - prices[j]
			}

			if j > 0 {
				currProfit += profits[j]
			}
			profits[i] = max(profits[i], currProfit)
		}
	}

	return profits[len(prices)-1]
}
