package buyandsellstock

import "testing"

func BuyAndSellStockCG(prices []int) int {
	maxGain := 0
	minNumber := 10_000
	for i := 0; i < len(prices); i++ {
		maxGain = max(maxGain, prices[i]-minNumber)
		minNumber = min(minNumber, prices[i])
	}
	return maxGain
}

func TestBuyAndSellStockCG(t *testing.T) {
	for i, test := range TestCases {
		actual := BuyAndSellStockCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %d, got: %d", i, test.output, actual)
		}
	}
}
