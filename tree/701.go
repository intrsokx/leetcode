package main

import "github.com/intrsokx/leetcode/model/treeModel"

//给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值
//都不同。
//
// 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,7,1,3], val = 5
//输出：[4,2,7,1,3,5]
//解释：另一个满足题目要求可以通过的树是：
//
//
//
// 示例 2：
//
//
//输入：root = [40,20,60,10,30,50,70], val = 25
//输出：[40,20,60,10,30,50,70,null,null,25]
//
//
// 示例 3：
//
//
//输入：root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
//输出：[4,2,7,1,3,5]
//
//
//
//
//
//
// 提示：
//
//
// 给定的树上的节点数介于 0 和 10^4 之间
// 每个节点都有一个唯一整数值，取值范围从 0 到 10^8
// -10^8 <= val <= 10^8
// 新值和原始二叉搜索树中的任意节点值都不同
//
// Related Topics 树
// 👍 189 👎 0

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
思路：分治法，找到最后一个满足条件的叶子节点插入即可
*/
func insertIntoBST(root *treeModel.TreeNode, val int) *treeModel.TreeNode {
	if root == nil {
		node := &treeModel.TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		return node
	}

	if root.Val.(int) > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)