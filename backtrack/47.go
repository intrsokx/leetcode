package main

import (
	"fmt"
	"sort"
)

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 数组 回溯
// 👍 748 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)

	backtrackV4(nums, list, visited, &ret)
	return ret
}

func backtrackV4(nums []int, list []int, visited []bool, ret *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*ret = append(*ret, ans)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		//上一个元素和当前相同，并且没有被访问过，则跳过
		//这句话反过来说，就是上一个元素和当前相同，并且被访问过了，那么当前元素就是下一个可选元素，则需要添加到访问路径中
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}

		list = append(list, nums[i])
		visited[i] = true
		backtrackV4(nums, list, visited, ret)
		list = list[:len(list)-1]
		visited[i] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 2}))
}
