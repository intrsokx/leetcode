package main

import "fmt"

//给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
// 假设你总是可以到达数组的最后一个位置。
//
//
//
// 示例 1:
//
//
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
//
// 示例 2:
//
//
//输入: [2,3,0,1,4]
//输出: 2
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 105
//
// Related Topics 贪心 数组 动态规划
// 👍 1051 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	//dp[i] 表示到i的最小跳数
	//dp[i] = min(dp[j]+1) 从小于i的j中找到最小的一个加一
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		//dp[i] 最大为i
		dp[i] = i
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min2num(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(nums)-1]
}

func min2num(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//贪心+动态
func jumpV2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		//取第一个能跳到的点
		idx := 0
		for idx < n && idx+nums[idx] < i {
			//说明从idx不能直接跳到i，所以要借助前一跳，才能跳到i
			idx++
		}

		//从idx跳一步可以跳到i
		dp[i] = dp[idx] + 1
	}

	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(jumpV2([]int{2, 3, 1, 1, 4}))
}
