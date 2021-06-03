package main
//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 递归
// 👍 888 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

//Definition for a binary tree node.
//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}

//分治法
func maxDepth(root *TreeNode) int {
	if root	== nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return 1 + left
	}
	return 1+right
}

//按照层次遍历的思路，每遍历完一层，就给count++
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	queue = append(queue, root)

	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}

		queue = queue[size:]
		count++
	}

	return count
}
//leetcode submit region end(Prohibit modification and deletion)
