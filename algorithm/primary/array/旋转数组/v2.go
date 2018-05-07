package main

import "fmt"

/**
	开辟一个新数组.然后搞一下.
	时间复杂度O(N).
	空间复杂度O(N).
 */
func rotateV2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	nLen := len(nums)
	a := make([]int, nLen, nLen)
	for i := range nums {
		a[(i+k)%nLen] = nums[i]
	}
	for i := range a {
		nums[i] = a[i]
	}
}

func main() {
	rotateV2([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
