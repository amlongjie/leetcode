package main

/**
	获取最大收益.一支股票可以多次买卖.
	那么当所有可能获利的情况都获利则获得的收益最高
	eg:7, 1, 5, 3, 6, 7,  4
	买:   1 3 6
	卖:   5 6 7
	利:   4 3 1
	时间复杂度O(N)
	空间复杂度O(1)
 */
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit([]int{1, 2, 3, 4, 5}))
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}
