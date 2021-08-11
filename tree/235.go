package main

//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树
// 👍 635 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//思路：因为是二叉搜索树，那么我们就可以根据根节点的值和p,q的值做比较来判断p,q是在左子树还是在右子树，或者分属左右子树
//递归
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	} else {
		return root
	}
}

//迭代
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	ret := root
	for {
		if ret.Val > p.Val && ret.Val > q.Val {
			ret = ret.Left
		} else if ret.Val < p.Val && ret.Val < q.Val {
			ret = ret.Right
		} else {
			return ret
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
