package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 你可以假设数组中无重复元素。
//
// 示例 1:
//
// 输入: [1,3,5,6], 5
//输出: 2
//
//
// 示例 2:
//
// 输入: [1,3,5,6], 2
//输出: 1
//
//
// 示例 3:
//
// 输入: [1,3,5,6], 7
//输出: 4
//
//
// 示例 4:
//
// 输入: [1,3,5,6], 0
//输出: 0
//
// Related Topics 数组 二分查找
// 👍 952 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1

	//循环退出条件start+1 == end
	for start+1 < end {
		mid := (end-start)/2 + start

		if nums[mid] <= target {
			start = mid
		} else {
			end = mid
		}
	}

	if nums[start] >= target {
		//>表示没有找到目标值，则start是目标值插入的位置，=表示找到
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		//nums[end] < target
		//target比所有值都大
		return end + 1
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
}
