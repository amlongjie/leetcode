package main

import "fmt"

/**
	不使用额外数据.
	通过移位来解决.
	时间复杂度O(N^2)
 */
func moveZeroesV1(nums []int) {
	start := 0
	end := len(nums)
	zC := 0
	for ; start < end-zC; {
		if nums[start] == 0 {
			preIndex := start
			for i := start + 1; i < end-zC; i++ {
				if nums[i] == 0 {
					continue
				}
				nums[i], nums[preIndex] = nums[preIndex], nums[i]
				preIndex = i
			}
			zC++
		}
		start++
	}
	fmt.Println(nums)
}

func main() {
	moveZeroesV1([]int{0})
	moveZeroesV1([]int{0, 1})
	moveZeroesV1([]int{0, 1, 2})
	moveZeroesV1([]int{0, 1, 0})
	moveZeroesV1([]int{0, 0, 1})
}
