package main

import "github.com/intrsokx/leetcode/model/treeModel"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层序遍历结果：
//
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 895 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type treeModel.TreeNode struct {
 *     Val int
 *     Left *treeModel.TreeNode
 *     Right *treeModel.TreeNode
 * }
 */

/**
思路：用一个队列记录一层的元素，然后扫描这一层元素添加下一层元素到队列（一个数进去出来一次，所以复杂度 O(logN)）
*/
func levelOrder(root *treeModel.TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*treeModel.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, l)

		for i := 0; i < l; i++ {
			node := queue[i]
			tmp[i] = node.Val.(int)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[l:]
	}

	return ret
}

func levelOrderV2(root *treeModel.TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*treeModel.TreeNode, 0)
	queue = append(queue, root)

	var tmp *treeModel.TreeNode
	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, l)

		for i := 0; i < l; i++ {
			tmp = queue[i]
			level = append(level, tmp.Val.(int))

			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}

		ret = append(ret, level)
		queue = queue[l:]
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
