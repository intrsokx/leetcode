package main

import "fmt"

//有n个物品和一个大小为m的背包， 给定数组A表示每个物品的大小，和数组V表示每个物品的价值。
//求最多能装入背包的总价值是多大？

//思路：dp[i][j]表示前i个物品，放入容量为j的背包中的最大价值是多少
func backpackII(m int, A, V []int) int {
	//dp[i][j]表示前i个物品，放入容量为j的背包中的最大价值是多少
	//dp[i][j] = dp[i-1][j] || dp[i-1][j-A[i]] + V[i]
	//dp[0][0] = 0
	//return dp[len[A]][m]

	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= m; j++ {
			//初始，不放当前物品
			dp[i][j] = dp[i-1][j]
			if j >= A[i-1] {
				//可以放，则需要选择是否要放当前物品
				dp[i][j] = max2numV5(dp[i][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}

	return dp[len(A)][m]
}

func max2numV5(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//m = 10
	//A = [2, 3, 5, 7]
	//V = [1, 5, 2, 4]
	fmt.Println(backpackII(10, []int{2, 3, 5, 7}, []int{1, 5, 2, 4}))

	//m = 10
	//A = [2, 3, 8]
	//V = [2, 5, 8]
	fmt.Println(backpackII(10, []int{2, 3, 8}, []int{2, 5, 8}))
}
