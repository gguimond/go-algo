package stockbuysell

func StockBuySell(prices []int) int {
	sell := 0
	buy := 0
	n := len(prices)
	gain := 0

	if n == 1 {
		panic("at least two prices")
	}

	for i := 0; i < n-1; {
		for i < n-1 && prices[i+1] <= prices[i] {
			i += 1
		}

		if i == n-1 {
			break
		}

		buy = i
		i += 1

		for i < n && prices[i] > prices[i-1] {
			i += 1
		}

		sell = i - 1
		gain += prices[sell] - prices[buy]
	}

	return gain
}

func main() {
}
