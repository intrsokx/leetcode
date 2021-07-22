package main

import "fmt"

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
//
//
// 示例 1：
//
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
//
// 示例 2：
//
//
//输入：digits = ""
//输出：[]
//
//
// 示例 3：
//
//
//输入：digits = "2"
//输出：["a","b","c"]
//
//
//
//
// 提示：
//
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
//
// Related Topics 哈希表 字符串 回溯
// 👍 1406 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

var tag map[byte][]byte

func letterCombinations(digits string) []string {
	tag = make(map[byte][]byte, 0)
	tag['2'] = []byte{'a', 'b', 'c'}
	tag['3'] = []byte{'d', 'e', 'f'}
	tag['4'] = []byte{'g', 'h', 'i'}
	tag['5'] = []byte{'j', 'k', 'l'}
	tag['6'] = []byte{'m', 'n', 'o'}
	tag['7'] = []byte{'p', 'q', 'r', 's'}
	tag['8'] = []byte{'t', 'u', 'v'}
	tag['9'] = []byte{'w', 'x', 'y', 'z'}
	if len(digits) == 0 {
		return []string{}
	}

	ret := make([]string, 0)
	list := make([]byte, 0)
	backtrackV6(digits, 0, list, &ret)
	return ret
}

func backtrackV6(digits string, pos int, list []byte, ret *[]string) {
	if len(list) == len(digits) {
		*ret = append(*ret, string(list))
		return
	}

	for i := pos; i < len(digits); i++ {
		for _, b := range tag[digits[i]] {
			list = append(list, b)
			backtrackV6(digits, i+1, list, ret)
			list = list[:len(list)-1]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
}
