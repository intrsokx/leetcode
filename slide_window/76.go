package main

import (
	"fmt"
	"math"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 105
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 字符串 滑动窗口
// 👍 1243 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	win := make(map[byte]int, 0)
	need := make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	//窗口
	left, right := 0, 0
	//匹配的次数
	match := 0
	start, end := 0, 0
	min := math.MaxInt32

	var c byte
	for right < len(s) {
		c = s[right]
		right++

		//如果当前字符是需要的字符，则加到窗口中
		if _, ok := need[c]; ok {
			win[c]++
			//如果窗口内当前字符数量等于所需数量，则match+1
			if win[c] == need[c] {
				match++
			}
		}

		//当所有字符数量都匹配之后，开始收缩窗口
		for match == len(need) {
			if right-left < min {
				min = right - left
				start = left
				end = right
			}

			c = s[left]
			left++
			if _, ok := need[c]; ok {
				//如果要排除的字符，是所需要的字符，并且移除该字符后，窗口内的字符不满足所需的字符，则match--
				//因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3; 所以在压死骆驼的最后一根稻草时，match才减一
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}

	if min == math.MaxInt32 {
		return ""
	}

	return s[start:end]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
