package main

import "fmt"

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 100
//
// Related Topics 数组 动态规划 矩阵
// 👍 924 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	//初始化dp数组，dp[i][j]表示从0,0开始到i,j的最小路径
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	//初始化dp[i][0] dp[0][j]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	//递推求解
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min2numV2(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func minPathSumV2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	//初始化dp
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min2numV2(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}

	return dp[n-1]
}

func min2numV2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
	fmt.Println(minPathSumV2(grid))
}
