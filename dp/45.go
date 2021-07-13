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
// 👍 1050 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//动态规划
func jump(nums []int) int {
	//dp[i] 表示从起点到当前位置的最小跳数
	//dp[i] = min(dp[j]+1, i )
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		//dp[i]的最大值为i
		dp[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				//dp[j]+1表示从初始节点跳到当前节点的跳数
				dp[i] = min2numV3(dp[j]+1, dp[i])
			}
		}
	}

	return dp[len(nums)-1]
}

func min2numV3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//贪心算法：正向查找可达的最大位置
//我们每次多找一步，每次找到下一步可到达的最远位置，这样就可以在线性时间内得到最少的跳跃次数
func jumpV2(nums []int) int {
	rightMost := 0
	step := 0
	end := 0

	for i := 0; i < len(nums)-1; i++ {
		rightMost = max2numV2(rightMost, nums[i]+i)

		if i == end {
			end = rightMost
			step++
		}
	}
	return step
}

func max2numV2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jumpV2([]int{2, 3, 1, 1, 4}))
}
