package main

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。
//
//
//示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
// Related Topics 数组 二分查找
// 👍 255 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	//循环退出条件left+1 == right
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	//结束后left+1 == right,则left是第一个大于target的索引

	return -1

	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)