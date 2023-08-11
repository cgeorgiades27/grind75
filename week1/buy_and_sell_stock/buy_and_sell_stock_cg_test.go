package buyandsellstock

import "testing"

func TestBuyAndSellStockCG(t *testing.T) {
	for i, test := range TestCases {
		actual := buyAndSellStockCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %d, got: %d", i, test.output, actual)
		}
	}
}
