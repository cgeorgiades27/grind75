package buyandsellstock

func buyAndSellStock(prices []int) int {
	lowestPrice, maxProfit := 10_000, 0
	for _, price := range prices {
		if price-lowestPrice > maxProfit {
			maxProfit = price - lowestPrice
		}
		if price < lowestPrice {
			lowestPrice = price
		}
	}

	return maxProfit
}
