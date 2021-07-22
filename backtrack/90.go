package main

import (
	"fmt"
	"sort"
)

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
//
//
//
// Related Topics 位运算 数组 回溯
// 👍 611 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	ret := make([][]int, 0)
	list := make([]int, 0)

	//排序是为了在选择的时候方便的进行剪枝操作
	sort.Ints(nums)
	backtrackV2(nums, 0, list, &ret)

	return ret
}

func backtrackV2(nums []int, pos int, list []int, ret *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*ret = append(*ret, ans)

	i := 0
	for i = pos; i < len(nums); i++ {
		//剪枝
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backtrackV2(nums, i+1, list, ret)
		list = list[:len(list)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	//[4,4,4,1,4]
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}
