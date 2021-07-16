package main

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
//
// Related Topics 树 二叉搜索树 动态规划 回溯 二叉树
// 👍 930 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	ans := make([]*TreeNode, 0)

	for i := start; i <= end; i++ {
		//以i为当前root节点，分别递归计算左右子树
		lefts := generate(start, i-1)
		rights := generate(i+1, end)

		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{
					Val:   i,
					Left:  lefts[j],
					Right: rights[k],
				}
				ans = append(ans, root)
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
