package main

import "fmt"

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//
// 返回符合要求的 最少分割次数 。
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：0
//
//
// 示例 3：
//
//
//输入：s = "ab"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
//
// Related Topics 字符串 动态规划
// 👍 444 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	//dp[i] 表示前i个字符组成的字符串最少需要分割的次数
	//trans: dp[i] = min(dp[j]+1) 满足j<i，并且[j+1,i]是回文子串
	//init: dp[i] = i-1， 前i个字符串最多需要分割i-1次，变成i个长度为1的回文子串
	dp := make([]int, len(s)+1)
	dp[0] = -1
	dp[1] = 0

	for i := 1; i <= len(s); i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			//j表示第几个元素，这里使用下标判断，所以要减一
			if isPalindrome(s, j+1-1, i-1) {
				dp[i] = min2numV2(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(s)]
}

func min2numV2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//isPalindrome
func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minCut("aab"))
}
