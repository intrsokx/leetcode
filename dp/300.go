package main

import "fmt"

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序
//列。
//
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -104 <= nums[i] <= 104
//
//
//
//
// 进阶：
//
//
// 你可以设计时间复杂度为 O(n2) 的解决方案吗？
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
// Related Topics 数组 二分查找 动态规划
// 👍 1696 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//求最长上升子序列：子序列跟子数组不一样，子序列可以不连续，子数组是连续的
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	//dp[i] 表示从0开始到i为止的最长上升子序列的长度
	//init: dp[i] = 1
	//trans: dp[i] = max(dp[j]+1) n[i]>n[j]

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max2numV2(dp[i], dp[j]+1)
			}
		}
	}

	max := dp[0]
	for _, ret := range dp {
		if ret > max {
			max = ret
		}
	}
	return max
}

func max2numV2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{3, 3, 3}))
}
