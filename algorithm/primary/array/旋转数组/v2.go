package main

/**
	循环移位.时间复杂度k * O(N)
	34个测试用例,过了33个.最后一个超时了.

 */
func rotateV2(nums []int, k int) {
	k = k % len(nums)
	a := nums
	for i := 0; i < k; i++ {
		nums = append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...)
	}
	copy(a, nums)
}

func main() {

}
