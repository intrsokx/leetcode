package main

import "fmt"

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
// 示例 4：
//
//
//输入：coins = [1], amount = 1
//输出：1
//
//
// 示例 5：
//
//
//输入：coins = [1], amount = 2
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 104
//
// Related Topics 广度优先搜索 数组 动态规划
// 👍 1365 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//动态规划：与其他动态规划不太一样，这里的i表示钱或者容量
func coinChange(coins []int, amount int) int {
	//dp[i] 表示能够组成金额i的硬币的最小个数
	//trans: dp[i] = min(dp[i-1], dp[i-2], dp[i-5])+1, i-coin >= 0
	//init: dp[i] = amount+1 初始化为最大值，求的是最小值，所以这块初始化为一个最大值
	//dp[0] = 0
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin >= 0 {
				//i-coin 表示组成金额i-coin需要多少个硬币，然后再加上需要当前coin这一枚硬币，刚好等于i
				dp[i] = min2numV4(dp[i], dp[i-coin]+1)
			}
		}
	}

	//结果等于初始最大值，则说明无法组成
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min2numV4(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 11))
}
