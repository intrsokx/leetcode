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
// 👍 1259 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//动态规划
func canJump(nums []int) bool {
	//dp[i]表示能否到达第i个位置
	//dp[i] = dp[j] && n[j]+j>=i 有没有一个小于i的j能跳到i

	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(nums)-1]
}

//贪心算法
//思路：遍历数组，每次寻找能走到的最远的距离
func canJumpV2(nums []int) bool {
	end := len(nums) - 1
	rightMost := 0

	if end <= rightMost {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if i <= rightMost {
			rightMost = max2num(rightMost, i+nums[i])
			if rightMost >= end {
				return true
			}
		}
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
	fmt.Println(canJumpV2([]int{0}))
}
