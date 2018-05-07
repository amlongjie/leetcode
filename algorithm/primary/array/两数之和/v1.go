package main

import "fmt"

/**
	使用Hash表.
	时间复杂度O(N).
	需要注意下面两个用例.
 */
func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, i}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSumV2([]int{3, 2, 4}, 6))
	fmt.Println(twoSumV2([]int{3, 3}, 6))
}
