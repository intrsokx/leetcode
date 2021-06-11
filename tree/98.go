package main

import (
	"fmt"
	"github.com/intrsokx/leetcode/model/treeModel"
)

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
 * type treeModel.TreeNode struct {
 *     Val int
 *     Left *treeModel.TreeNode
 *     Right *treeModel.TreeNode
 * }
 */

/**
思路一：中序遍历，然后判断结果是否有序
思路二：分治法，判断左max < 根 < 右min
*/
func isValidBST(root *treeModel.TreeNode) bool {
	return true
}

func isValidBST_v1(root *treeModel.TreeNode) bool {
	lst := make([]int, 0)
	inOrder(root, &lst)

	for i := 0; i < len(lst)-1; i++ {
		if lst[i] >= lst[i+1] {
			return false
		}
	}

	return true
}
func inOrder(node *treeModel.TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, ret)
	*ret = append(*ret, node.Val.(int))
	inOrder(node.Right, ret)
}

type ResultType struct {
	IsValid bool
	Max     *treeModel.TreeNode
	Min     *treeModel.TreeNode
}

func isValidBST_v2(root *treeModel.TreeNode) bool {
	ret := helper1(root)
	return ret.IsValid
}

//左max < 根 < 右min
func helper1(node *treeModel.TreeNode) ResultType {
	ret := ResultType{}

	if node == nil {
		ret.IsValid = true
		return ret
	}

	left := helper1(node.Left)
	right := helper1(node.Right)

	if !left.IsValid || !right.IsValid {
		ret.IsValid = false
		return ret
	}
	if left.Max != nil && left.Max.Val.(int) >= node.Val.(int) {
		ret.IsValid = false
		return ret
	}
	if right.Min != nil && right.Min.Val.(int) <= node.Val.(int) {
		ret.IsValid = false
		return ret
	}

	//当前子树是二叉搜索树
	ret.IsValid = true
	ret.Min = node
	if left.Min != nil {
		ret.Min = left.Min
	}
	ret.Max = node
	if right.Max != nil {
		ret.Max = right.Max
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	ret := make([]int, 0)
	inOrder(nil, &ret)

	fmt.Println(ret)

	var a, b *treeModel.TreeNode
	a = treeModel.GenerateNode(2, nil, nil)
	b = treeModel.GenerateNode(2, nil, nil)

	root := treeModel.GenerateNode(2, a, b)

	fmt.Println(isValidBST_v1(root))
}
