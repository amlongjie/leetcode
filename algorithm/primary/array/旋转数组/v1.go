package main

/**
	初始: 1,2,3,4,5,6. k = 3
	目标: 4,5,6, 1,2,3
	有点像全部反转: 6,5,4,3,2,1 => 4,5,6,3,2,1 => 4,5,6,1,2,3
	时间复杂度O(N)
	空间复杂度O(1)
 */
func rotate(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
func reverse(ints []int, x int, y int) {
	for x < y {
		ints[x], ints[y] = ints[y], ints[x]
		x++
		y--
	}
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

}
