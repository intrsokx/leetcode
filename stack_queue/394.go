package main

import (
	"fmt"
	"strconv"
)

//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
//
//
// 示例 1：
//
// 输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//
// 示例 2：
//
// 输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
//
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
//
// Related Topics 栈 递归 字符串
// 👍 795 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			//得到要重复的字节数组的翻转结果
			temp := make([]byte, 0)
			for len(stack) >= 0 && stack[len(stack)-1] != '[' {
				ch := stack[len(stack)-1]
				temp = append(temp, ch)
				stack = stack[:len(stack)-1]
			}
			//'[' pop
			stack = stack[:len(stack)-1]

			//calc num
			idx := 1
			for len(stack)-idx >= 0 && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]

			//write reverse(temp) n times
			n, _ := strconv.Atoi(string(num))
			for j := 0; j < n; j++ {
				for j := len(temp) - 1; j >= 0; j-- {
					stack = append(stack, temp[j])
				}
			}

		default:
			ch := s[i]
			stack = append(stack, ch)
		}
	}

	return string(stack)

	//range 遍历的是字符rune，在本题中需要的是byte(uint8)
	//for _, v := range s {
	//	if v == ']' {
	//
	//	} else {
	//		stack = append(stack, v)
	//	}
	//}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
}
