package main

import "fmt"

/**
	按照算数加法的方式计算即可.
	时间:O(N)
 */
func plusOne(digits []int) []int {
	carry := 1
	sLen := len(digits)
	res := make([]int, sLen+1, sLen+1)
	for i := sLen - 1; i >= 0; i-- {
		tmp := digits[i] + carry
		if tmp >= 10 {
			carry = 1
			res[i+1] = tmp - 10
		} else {
			carry = 0
			res[i+1] = tmp
		}
	}
	if carry == 1 {
		res[0] = 1
		return res
	} else {
		return res[1:]
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
}
