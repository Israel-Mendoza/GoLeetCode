package topinterview150

import "math"

func MaxProfit(prices []int) int {

	minStockPrice := math.MaxInt
	maxStockProfit := 0

	for _, stockPrice := range prices {
		minStockPrice = min(minStockPrice, stockPrice)
		maxStockProfit = max(stockPrice-minStockPrice, maxStockProfit)
	}

	return maxStockProfit
}
