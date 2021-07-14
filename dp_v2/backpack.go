package main

import "fmt"

//在n个物品中挑选若干物品装入背包，最多能装多满
//eg: A = [3, 4, 8, 5], m=10   ret=9

func backPack(m int, A []int) int {
	// write your code here
	// f[i][j] 前i个物品，是否能装j
	// f[i][j] =f[i-1][j] f[i-1][j-a[i] j>a[i]
	// f[0][0]=true f[...][0]=true
	// f[n][X]
	f := make([][]bool, len(A)+1)
	for i := 0; i <= len(A); i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			//如果不放第i件物品
			f[i][j] = f[i-1][j]
			if j-A[i-1] >= 0 && f[i-1][j-A[i-1]] {
				//如果放第i件物品
				f[i][j] = true
			}
		}
	}
	for i := m; i >= 0; i-- {
		if f[len(A)][i] {
			return i
		}
	}
	return 0
}

func backPackV2(m int, A []int) int {
	//dp[i][j] 表示第i个物品，在背包容量为j的情况下，最大的装载量
	//dp[i][j] = max(dp[i-1][j], dp[i-1][j-A[i-1]] + A[i-1]

	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 0

	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= A[i-1] {
				dp[i][j] = max2numV4(dp[i-1][j], dp[i-1][j-A[i-1]]+A[i-1])
			}
		}
	}

	return dp[len(A)][m]
}

func max2numV4(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(backPack(4, []int{2, 3}))
	fmt.Println(backPackV2(4, []int{2, 3}))

	fmt.Println(backPack(10, []int{3, 4, 8, 5}))
	fmt.Println(backPackV2(10, []int{3, 4, 8, 5}))
}
