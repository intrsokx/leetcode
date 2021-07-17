package main

import "fmt"

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 异位词 指字母相同，但排列不同的字符串。
//
//
//
// 示例 1:
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 104
// s 和 p 仅包含小写字母
//
// Related Topics 哈希表 字符串 滑动窗口
// 👍 564 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	win := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	match := 0

	ret := make([]int, 0)

	var ch byte
	for right < len(s) {
		ch = s[right]
		right++

		if _, ok := need[ch]; ok {
			win[ch]++
			if win[ch] == need[ch] {
				match++
			}
		}

		//窗口大小大于等于所需大小，判断是否找到结果 并 收缩窗口
		if right-left >= len(p) {
			if match == len(need) {
				ret = append(ret, left)
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				//收缩之后，就不满足need条件了
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
