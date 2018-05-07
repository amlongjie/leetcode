package main

import "fmt"

/**
	用一个新数组.
	时间:O(N)
	空间:O(N)

 */
func moveZeroes(nums []int) {
	start := 0
	end := len(nums) - 1
	tmp := make([]int, len(nums), len(nums))
	for _, v := range nums {
		if v == 0 {
			tmp[end] = v
			end--
		} else {
			tmp[start] = v
			start++
		}
	}
	for i := range tmp {
		nums[i] = tmp[i]
	}
	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{1, 2, 0, 0, 3, 4})
}
