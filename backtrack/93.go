package main

import (
	"fmt"
	"strconv"
)

//给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
//和 "192.168@1.1" 是 无效 IP 地址。
//
//
//
// 示例 1：
//
//
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
//
//
// 示例 2：
//
//
//输入：s = "0000"
//输出：["0.0.0.0"]
//
//
// 示例 3：
//
//
//输入：s = "1111"
//输出：["1.1.1.1"]
//
//
// 示例 4：
//
//
//输入：s = "010010"
//输出：["0.10.0.10","0.100.1.0"]
//
//
// 示例 5：
//
//
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3000
// s 仅由数字组成
//
// Related Topics 字符串 回溯
// 👍 625 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	list := make([]string, 0)

	//根据IP地址的特点进行过滤，否则会超时
	if len(s) < 4 || len(s) > 12 {
		return ret
	}

	backtrackV8(s, 0, list, &ret)
	return ret
}

//pos 可选择子串的起始下标
//list 选择的结果
//ret 结果
func backtrackV8(s string, pos int, list []string, ret *[]string) {
	//满足条件，将选择的结果加入到最终结果中
	if pos == len(s) && len(list) == 4 {
		ip := ""
		for i := 0; i < 4; i++ {
			if i == 3 {
				ip += list[i]
			} else {
				ip += list[i] + "."
			}
		}
		*ret = append(*ret, ip)
		return
	} else if pos == len(s) {
		//字符串选完了，但是没有找到正确的路径（list）
		return
	}

	//IP地址的先决条件，以0开头的只能是0
	if s[pos] == '0' {
		list = append(list, "0")
		backtrackV8(s, pos+1, list, ret)
		list = list[:len(list)-1]
		return
	}

	//选择s[pos, i) 看其是否满足条件
	for i := pos + 1; i <= len(s); i++ {
		//s[start:i]组成的数字是否满足IP规范
		if !isValid(s[pos:i]) {
			break
		}

		list = append(list, s[pos:i])
		backtrackV8(s, i, list, ret)
		list = list[:len(list)-1]
	}
}

func isValid(s string) bool {
	n, e := strconv.Atoi(s)
	if e != nil {
		return false
	}
	if n > 0 && n < 256 {
		return true
	}

	return false
	//if s > "0" && s < "256" {
	//	return true
	//}
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	s1 := "0"
	s2 := "255"

	fmt.Println("1" > s1)
	fmt.Println("1" < s2)

	fmt.Println("222" > s1)
	fmt.Println("122" < s2)

	fmt.Println("0" > s1)
	fmt.Println("256" < s2)

	fmt.Println(restoreIpAddresses("25525511135"))
}
