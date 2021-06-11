package main

import "fmt"

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
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
// 返回其自底向上的层序遍历为：
//
//
//[
//  [15,7],
//  [9,20],
//  [3]
//]
//
// Related Topics 树 广度优先搜索
// 👍 446 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := make([]*TreeNode, 0)
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

	reverseByUpDown(ret)
	return ret
}

//对二维切片上下翻转
func reverseByUpDown(nums [][]int) {
	rows := len(nums)

	for i := 0; i < rows/2; i++ {
		nums[i], nums[rows-i-1] = nums[rows-i-1], nums[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5},
		[]int{6, 7, 8, 9},
		[]int{1, 2, 3},
	}

	reverseByUpDown(nums)
	fmt.Println(nums)
}
