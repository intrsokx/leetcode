package main

import "fmt"

//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
//
//
//
// 示例 1：
//
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 104
// s1 和 s2 仅包含小写字母
//
// Related Topics 哈希表 双指针 字符串 滑动窗口
// 👍 377 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	win := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	match := 0

	var ch byte
	for right < len(s2) {
		ch = s2[right]
		right++

		if _, ok := need[ch]; ok {
			win[ch]++
			if win[ch] == need[ch] {
				match++
			}
		}

		//若当前窗口大小大于等于所需字符串的长度，则判断是否match成功，并收缩窗口
		if right-left >= len(s1) {
			//如果匹配成功，则返回
			if match == len(need) {
				return true
			}

			//收缩窗口，并判断match
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))

	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
