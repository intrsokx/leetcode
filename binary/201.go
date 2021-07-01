package main

//给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
//。
//
//
//
// 示例 1：
//
//
//输入：left = 5, right = 7
//输出：4
//
//
// 示例 2：
//
//
//输入：left = 0, right = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：left = 1, right = 2147483647
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= left <= right <= 231 - 1
//
// Related Topics 位运算
// 👍 297 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//思路：left...right 多个连续数字按位与
//根据与操作的定义，只要其中有一个0则结果为0，只有都为1，结果才为1
//8~11
//1000
//1001
//1010
//1011
//结果 10,
//因为是多个连续数字，所以相邻的两个数字之间的最后一位不是0就是1，即本题变成了求多个连续数字相等的前缀是多少，
func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		count++
	}

	// left==right
	//left<<count 求相等前缀时右移了多少位，现在结果后面就补多少个0
	return left << count
}

//leetcode submit region end(Prohibit modification and deletion)
