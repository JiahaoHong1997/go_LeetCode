package ci63

func maxProfit(prices []int) int {

	if len(prices) == 0 {
        return 0
    }
	minPrice := prices[0]
	ret := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			continue
		}
		ret = max(ret, prices[i]-minPrice)
	}
	return ret
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}