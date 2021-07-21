package main

import "fmt"

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics 数组 回溯
// 👍 1455 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//思路：需要记录已经选择过的元素
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	list := make([]int, 0)
	visited := make(map[int]bool, len(nums))

	backtrackV3(nums, list, visited, &ret)
	return ret
}

func backtrackV3(nums []int, list []int, visited map[int]bool, ret *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*ret = append(*ret, ans)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		}
		list = append(list, nums[i])
		visited[nums[i]] = true

		backtrackV3(nums, list, visited, ret)

		list = list[:len(list)-1]
		visited[nums[i]] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
