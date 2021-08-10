package main

import (
	"fmt"
	"strings"
)

//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
//
//string convert(string s, int numRows);
//
//
//
// 示例 1：
//
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//
//示例 2：
//
//
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//
//
// 示例 3：
//
//
//输入：s = "A", numRows = 1
//输出："A"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
//
// Related Topics 字符串
// 👍 1230 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	ret := make([]string, numRows)
	//当s的长度小于行数，还没有走到最后一行，字符串就用完了
	//如果只有一行，直接输出字符串
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	//ret的下标
	idx := 0
	//标记位，用来对idx加加或减减
	tag := -1
	for i := 0; i < len(s); i++ {
		//走到了第一行或者最后一行，需要转换标记位
		if idx == 0 || idx == numRows-1 {
			tag = -tag
		}

		ret[idx] += string(s[i])
		idx = idx + tag
	}

	return strings.Join(ret, "")
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(convert("AB", 1))
}
