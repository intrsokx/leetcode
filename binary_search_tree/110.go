package main

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
//
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
//
//
// 示例 3：
//
//
//输入：root = []
//输出：true
//
//
//
//
// 提示：
//
//
// 树中的节点数在范围 [0, 5000] 内
// -104 <= Node.val <= 104
//
// Related Topics 树 深度优先搜索 递归
// 👍 695 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ResultTypeV2 struct {
	Valid  bool
	Height int
}

func isBalanced(root *TreeNode) bool {
	return dfsV2(root).Valid
}

func dfsV2(root *TreeNode) *ResultTypeV2 {
	if root == nil {
		return &ResultTypeV2{Valid: true, Height: 0}
	}
	ret := &ResultTypeV2{
		Valid:  false,
		Height: 0,
	}

	left := dfsV2(root.Left)
	right := dfsV2(root.Right)

	if left.Valid && right.Valid && abs(left.Height-right.Height) <= 1 {
		ret.Valid = true
	}
	ret.Height = max2numV2(left.Height, right.Height) + 1

	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max2numV2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
