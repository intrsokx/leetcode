package main

import "fmt"

//格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//
// 给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。
//
// 格雷编码序列必须以 0 开头。
//
//
//
// 示例 1:
//
// 输入: 2
//输出: [0,1,3,2]
//解释:
//00 - 0
//01 - 1
//11 - 3
//10 - 2
//
//对于给定的 n，其格雷编码序列并不唯一。
//例如，[0,2,3,1] 也是一个有效的格雷编码序列。
//
//00 - 0
//10 - 2
//11 - 3
//01 - 1
//
// 示例 2:
//
// 输入: 0
//输出: [0]
//解释: 我们定义格雷编码序列必须以 0 开头。
//     给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
//     因此，当 n = 0 时，其格雷编码序列为 [0]。
//
// Related Topics 位运算 数学 回溯
// 👍 321 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func grayCode(n int) []int {
	ret := make([]int, 0)
	ret = append(ret, 0)

	//每次要增加的增量（在之前的逆序结果上的最高位 加1）
	//00
	//01
	//11
	//10
	//110
	//111
	//101
	//100
	mark := 1
	for i := 1; i <= n; i++ {
		l := len(ret)
		//逆序
		for j := l - 1; j >= 0; j-- {
			ret = append(ret, ret[j]+mark)
		}
		//最高位1
		mark = mark << 1
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(0))
}
