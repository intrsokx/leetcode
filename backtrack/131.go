package main

import "fmt"

//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：[["a"]]
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 16
// s 仅由小写英文字母组成
//
// Related Topics 字符串 动态规划 回溯
// 👍 778 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//分析：回溯结束条件：s为""
//从其实位置判断s[pos,i]是否是回文子串，若是，则选择
func partition(s string) [][]string {
	ret := make([][]string, 0)
	list := make([]string, 0)

	backtrackV7(s, 0, list, &ret)
	return ret
}

func backtrackV7(s string, pos int, list []string, ret *[][]string) {
	if pos == len(s) {
		ans := make([]string, len(list))
		copy(ans, list)
		*ret = append(*ret, ans)
		return
	}

	for i := pos + 1; i <= len(s); i++ {
		if !isPalindrome(s[pos:i]) {
			continue
		}

		list = append(list, s[pos:i])
		backtrackV7(s, i, list, ret)
		list = list[:len(list)-1]
	}
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(partition("abb"))
	fmt.Println(partition("cdd"))
}
