package main

import "fmt"

/**
	遍历所有情况.O(N^2)
 */
func twoSum(nums []int, target int) []int {
	sLen := len(nums)
	for i := 0; i < sLen-1; i++ {
		for j := i + 1; j < sLen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
