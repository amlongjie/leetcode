package main

/**
排序数组去重,遍历一遍
时间复杂度:O(N)
空间复杂度:O(1)
 */
func removeDuplicates(nums []int) int {
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[res] = nums[i]
		res++
	}
	return res
}

func main() {
	println(removeDuplicates([]int{1, 2, 3, 4}))
	println(removeDuplicates([]int{1, 1, 2}))
	println(removeDuplicates([]int{1, 1, 2, 2, 3}))
	println(removeDuplicates([]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 5}))
}
