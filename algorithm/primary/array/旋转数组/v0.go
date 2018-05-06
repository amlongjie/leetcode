package main

/**
	循环移位.时间复杂度k * O(N)
	34个测试用例,过了33个.最后一个超时了.

 */
func rotateV0(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}
	if k == 0 {
		return
	}
	for i := 0; i < k; i++ {
		tmp := nums[0]
		for j := 0; j < len(nums); j++ {
			nums[(j+1)%len(nums)], tmp = tmp, nums[(j+1)%len(nums)]
		}
	}
}

func main() {

}
