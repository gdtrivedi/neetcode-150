func maxProfit(prices []int) int {
    maxProfit := 0
	minValue := math.MaxInt

	for _, p := range prices {
		if p < minValue {
			minValue = p
		}

		if p-minValue > maxProfit {
			maxProfit = p - minValue
		}
	}

	return maxProfit
}
