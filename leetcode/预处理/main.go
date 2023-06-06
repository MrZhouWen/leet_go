package 预处理

/**
给出mxn的矩阵，元素值 为0/1
题目:找出矩阵中【元素值为1，且其所在行和列的其他位置全是0的】的总个数

思路：预处理每一行和每一列的和，这样就可以将总的时间复杂度 控制在 mxn
T = O(mxn) S= O(m+n)预处理需要空间存储
*/

func numSpecial(mat [][]int) (ans int) {
	sumRow := make([]int, len(mat))
	sumCol := make([]int, len(mat[0]))
	for i, row := range mat {
		for j, x := range row {
			sumRow[i] += x
			sumCol[j] += x
		}
	}
	for i, row := range mat {
		for j, x := range row {
			if x == 1 && sumRow[i] == 1 && sumCol[j] == 1 {
				ans++
			}
		}
	}
	return ans
}
