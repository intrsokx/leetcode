package main

//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回锯齿形层序遍历如下：
//
//
//[
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
// 👍 452 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isReverse := false
	for len(queue) > 0 {
		l := len(queue)
		lst := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[i]
			lst = append(lst, node.Val.(int))
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]

		if isReverse {
			reverse(lst)
		}
		isReverse = !isReverse
		ret = append(ret, lst)
	}

	return ret
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
