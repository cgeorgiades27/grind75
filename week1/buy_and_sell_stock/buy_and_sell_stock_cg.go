package buyandsellstock

func buyAndSellStockCG(prices []int) int {
	maxGain := 0
	minNumber := 10_000
	for i := 0; i < len(prices); i++ {
		maxGain = max(maxGain, prices[i]-minNumber)
		minNumber = min(minNumber, prices[i])
	}
	return maxGain
}
