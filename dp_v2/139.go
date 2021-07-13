package main

import "fmt"

//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
// 说明：
//
//
// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
//
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//
//
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//     注意你可以重复使用字典中的单词。
//
//
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false
//
// Related Topics 字典树 记忆化搜索 哈希表 字符串 动态规划
// 👍 1049 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	//dp[i] 表示前i个字符能否被分割
	//trans: dp[i] = dp[j] && n[j+1,i] in dict
	//dp[0] = true

	dict := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			//j, i 是第几个字符
			//dp[j] 表示前j个字符是否能被分割
			if dp[j] && dict[s[j+1-1:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	dict := map[string]bool{
		"hello": true,
		"world": true,
	}

	fmt.Println(dict["hello"])
	fmt.Println(dict["aaa"])

	fmt.Println(wordBreak("helloworl", []string{"hello", "world"}))
}

/**
序列类型动态规划 小结：
1、常见处理方式是给 0 位置占位，这样处理问题时可以一视同仁，初始化则在原来的 基础上length+1，返回结果f[n]
2、状态可以为前i个
3、初始化length+1
4、取值 index=i-1
5、返回值 f[n] 或者 f[m][n]
*/
