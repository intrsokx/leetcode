package main

import "fmt"

//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 104
// 0 <= nums[i] <= 105
//
// Related Topics 贪心 数组 动态规划
// 👍 1258 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	//dp[i] 表示从初始位置是否能跳到下标为i的位置

	dp := make([]bool, len(nums))
	dp[0] = true

	//dp[i] = dp[j] 有没有小于i的一个j能跳到i
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			//dp[j]是可达的 并且能从j 跳到 i
			if dp[j] && nums[j] >= (i-j) {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}

//贪心算法：寻找能跳的最远的位置，如果大于等于最后一个位置则可达
func canJumpV2(nums []int) bool {
	rightMost := 0
	last := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = max2num(rightMost, i+nums[i])
		}
	}
	if rightMost >= last {
		return true
	}
	return false
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
