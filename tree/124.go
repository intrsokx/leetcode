package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
给定一个非空二叉树，返回其最大路径和。

思路：分治法，分为三种情况：1、左子树最大路径和最大，2、右子树最大路径和最大，3、左边+右边+根节点最大，
需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即可
*/
func maxPathSum(root *TreeNode) int {
	ret := helper(root)

	return ret.MaxPath
}

type Result struct {
	SinglePath int
	MaxPath    int
}

func helper(node *TreeNode) Result {
	if node == nil {
		return Result{
			SinglePath: 0,
			MaxPath:    math.MinInt64,
		}
	}

	//divide
	left := helper(node.Left)
	right := helper(node.Right)

	//conquer
	ret := Result{}
	//单边最大值
	if left.SinglePath > right.SinglePath {
		ret.SinglePath = max2num(left.SinglePath+node.Val.(int), 0)
	} else {
		ret.SinglePath = max2num(right.SinglePath+node.Val.(int), 0)
	}

	//左边 + 右边 + 根 的最大值
	maxPath := max2num(left.MaxPath, right.MaxPath)
	ret.MaxPath = max2num(maxPath, left.SinglePath+right.SinglePath+node.Val.(int))

	return ret
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
