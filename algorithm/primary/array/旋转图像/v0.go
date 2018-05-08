package main

/**
	题目要求原地修改.

	先上下翻转.
	再根据轴对称翻转.
	eg:
	1 2 3    7 8 9		7 4 1
	4 5 6 => 4 5 6 =>  	8 5 2 (针对753那条线对称)
	7 8 9    1 2 3		9 6 3
 */
func rotate(matrix [][]int) {
	mLen := len(matrix)
	for i := 0; i < mLen/2; i++ {
		for j := 0; j < mLen; j++ {
			matrix[i][j], matrix[mLen-i-1][j] = matrix[mLen-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < mLen; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	rotate([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
}
