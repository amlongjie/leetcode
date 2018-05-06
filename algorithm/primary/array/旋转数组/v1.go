package main

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
