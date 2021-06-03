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

/**
思路：分治法，左边平衡 && 右边平衡 && 左右两边高度 <= 1，
因为需要返回是否平衡及高度，要么返回两个数据，要么合并两个数据，所以用-1 表示不平衡，>0 表示树高度（二义性：一个变量有两种含义）

tips:一般在工程实践中，不建议使用一个变量表示两种含义，推荐使用两个变量返回具体的含义

不能计算一个最大高度跟最小高度之差来判断一棵树是否是平衡的。
因为高度平衡二元树定义：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 */
func maxDepthAndIsBalanced(node *TreeNode) (depth int, isBalance bool) {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	if node == nil {
		return 0, true
	}

	height1, ok1 := maxDepthAndIsBalanced(node.Left)
	height2, ok2 := maxDepthAndIsBalanced(node.Right)

	if !ok1	|| !ok2 || abs(height1-height2) > 1 {
		return 0, false
	}

	if height1 > height2 {
		return 1+height1, true
	}
	return 1+height2, true
}

func isBalanced(root *TreeNode) bool {
	var abs func(num int) int
	var maxDepth func(node *TreeNode) int

	abs = func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := maxDepth(node.Left), maxDepth(node.Right)
		//这块因为返回的变量具有二义性，需做个判断
		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1
		}

		if left > right {
			return 1+left
		}
		return 1+right
	}

	depth := maxDepth(root)
	if depth == -1 {
		return false
	}
	return true
}

func isBalanced1(root *TreeNode) bool {
	_, ok := maxDepthAndIsBalanced(root)
	return ok
}
//leetcode submit region end(Prohibit modification and deletion)
