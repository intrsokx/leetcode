package main

import "math"

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
//
// Related Topics 树 深度优先搜索 递归
// 👍 1092 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ResultType struct {
	IsValid bool
	Max     int
	Min     int
}

func isValidBST(root *TreeNode) bool {
	return dfs(root).IsValid
}

func dfs(root *TreeNode) *ResultType {
	ret := &ResultType{}
	if root == nil {
		ret.IsValid = true
		ret.Max = math.MinInt32
		ret.Min = math.MaxInt32
		return ret
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	//当期那节点满足排序二叉树的条件
	//左边的最大值 小于 当前值 小于 右边的最大值
	if left.IsValid && right.IsValid && root.Val >= left.Max && root.Val <= right.Min {
		ret.IsValid = true
	}

	//计算当前节点的最大最小值
	ret.Max = max3num(left.Max, root.Val, right.Max)
	ret.Min = min3num(left.Min, root.Val, right.Min)

	return ret
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func max3num(a, b, c int) int {
	return max2num(max2num(a, b), c)
}

func min2num(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min3num(a, b, c int) int {
	return min2num(min2num(a, b), c)
}

//leetcode submit region end(Prohibit modification and deletion)

//[5,1,4,null,null,3,6]
