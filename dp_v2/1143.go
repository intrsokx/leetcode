package main

//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//
//
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
//
//
// 示例 1：
//
//
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
//
// 示例 2：
//
//
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc" ，它的长度为 3 。
//
//
// 示例 3：
//
//
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= text1.length, text2.length <= 1000
// text1 和 text2 仅由小写英文字符组成。
//
// Related Topics 字符串 动态规划
// 👍 604 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//求最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j] 表示s1的前i个子串 与 s2的前j个子串 的最长公共子序列的长度
	//init: dp[0][0] = 0
	//trans: dp[i][j] = dp[i-1][j-1]+1 (s1[i]==s2[j])
	//下标idx: i-1, j-1

	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//相等取左上值加1，否则，取左边或上边最大的值
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max2numV3(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max2numV3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
