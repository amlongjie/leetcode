package main

import "fmt"

/**
	如果不原地修改.可以使用扩展内存.
	TODO
 */
func rotateV1(matrix [][]int) {
	mLen := len(matrix)
	tmpMatrix := make([][]int, mLen, mLen)
	for i := 0; i < mLen/2; i++ {
		oX, oY := i, i
		tX, tY := i, mLen-1-i
		for j := oX; j <= tY; j++ {
			tmpMatrix[tX][tY] = matrix[oX][oY]
		}
	}
	fmt.Println(tmpMatrix)
}

func main() {
	rotateV1([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}
