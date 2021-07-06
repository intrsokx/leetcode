package main

import "fmt"

//已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变
//化后可能得到：
//
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
// 若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
//
//
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2],
//..., a[n-2]] 。
//
// 给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,5]
//输出：1
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,0,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 5000
// -5000 <= nums[i] <= 5000
// nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
//
//
//
//
// 进阶：
//
//
// 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
// 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
//
// Related Topics 数组 二分查找
// 👍 378 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//思路：跳过重复元素，mid值和end值比较
func findMinV2(nums []int) int {
	start, end := 0, len(nums)-1

	//for quit: start+1 == end
	for start+1 < end {
		//skip repeat item
		//这里因为start < end， 所以end-1, start+1一定满足数组下标条件
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}

		mid := (end-start)/2 + start
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}

	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findMinV2([]int{2, 2, 2, 0, 1}))
}
