package leet

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mp := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			if price-minPrice > mp {
				mp = price - minPrice
			}
		}
	}

	return mp
}
