package main

import "fmt"

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// Related Topics 位运算 数组 回溯
// 👍 1250 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &ret)
	return ret
}

//nums: 初始数组
//pos: 遍历下标
//list: 临时结果
//ret: 结果
func backtrack(nums []int, pos int, list []int, ret *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*ret = append(*ret, ans)

	var i int
	for i = pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, ret)
		list = list[:len(list)-1]
	}

	//这里加个断点，方便调试
	fmt.Println("will return")
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(subsets([]int{1, 2}))
}
