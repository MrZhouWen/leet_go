package 动态规划

/**
题目：给定素因子[3， 5，7], 找出不包含其他素因子的第k个数
例如：前几个数 1， 3， 5，7， 9 ， 15， 21

思路：3指针（dp）
定义数组dp, dp[i]表示第i个数，第k个即为dp[k], dp[1] = 1
定义3个指针，p3, p5, p7分别指向最小的数的下标，初始都为1; 令
	x3 = dp[p3]*3
	x5 = dp[p5]*5
	x7 = dp[p7]*7
当 2 <= i <= k时， 令dp[i] = min(x3, x5, x7), 被选中的最小的数，将其下标+1即可，比如 向x5 是3者最小的，p5++

T = O(k) S= O(k)
*/
func getKthMagicNumber(k int) int {
	dp := make([]int, k+1) //下标从1开始的
	dp[1] = 1
	p3, p5, p7 := 1, 1, 1
	for i := 2; i <= k; i++ {
		x3, x5, x7 := dp[p3]*3, dp[p5]*5, dp[p7]*7
		dp[i] = min(x3, min(x5, x7))
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
		if dp[i] == x7 {
			p7++
		}
	}
	return dp[k]
}
