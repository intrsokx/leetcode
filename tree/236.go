package main

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
//
//
// 示例 1：
//
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出：3
//解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
//
//
// 示例 2：
//
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出：5
//解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
//
//
// 示例 3：
//
//
//输入：root = [1,2], p = 1, q = 2
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [2, 105] 内。
// -109 <= Node.val <= 109
// 所有 Node.val 互不相同 。
// p != q
// p 和 q 均存在于给定的二叉树中。
//
// Related Topics 树
// 👍 1146 👎 0

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
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

分治法：自顶向下分解问题，若当前节点的左子树或者右子树是两个节点的共同祖先，则返回子树节点,否则返回当前节点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//condition judge 自底向顶计算答案，如果p q中有一个跟root相等，那么说明离其最近的公共祖先就可能是root本身，
	//需要由其上一层去判断是不是当前root
	if p == root || q == root {
		return root
	}

	//divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//若左右两边都不为空，则当前根节点为祖先
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil { //这里不直接返回right，是因为这样写表意比较明确
		return right
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)

var parent map[*TreeNode]*TreeNode

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		parent[node.Left] = node
		dfs(node.Left)
	}
	if node.Right != nil {
		parent[node.Right] = node
		dfs(node.Right)
	}
}

//先根据子节点p，不断的向上查找其父节点，直到空为止，边遍历边将此路径做标记，这时我们就得到了一条从p到根节点的路径；
//然后从q节点向上遍历，若两个节点有有公共父节点，那么在这条向上遍历的路径上就一定会有交点，则第一次遇到的交点为最近的公共父节点。
func lowestCommonAncestor0(root, p, q *TreeNode) *TreeNode {
	parent = make(map[*TreeNode]*TreeNode)
	if root == nil {
		return nil
	}
	parent[root] = nil

	dfs(root)

	visited := make(map[*TreeNode]bool)
	for p != nil {
		visited[p] = true
		p = parent[p]
	}
	for q != nil {
		if visited[q] {
			return q
		}
		q = parent[q]
	}
	return nil
}
