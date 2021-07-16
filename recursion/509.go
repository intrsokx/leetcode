package main

import "fmt"

//斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
//
//
//F(0) = 0，F(1) = 1
//F(n) = F(n - 1) + F(n - 2)，其中 n > 1
//
//
// 给你 n ，请计算 F(n) 。
//
//
//
// 示例 1：
//
//
//输入：2
//输出：1
//解释：F(2) = F(1) + F(0) = 1 + 0 = 1
//
//
// 示例 2：
//
//
//输入：3
//输出：2
//解释：F(3) = F(2) + F(1) = 1 + 1 = 2
//
//
// 示例 3：
//
//
//输入：4
//输出：3
//解释：F(4) = F(3) + F(2) = 2 + 1 = 3
//
//
//
//
// 提示：
//
//
// 0 <= n <= 30
//
// Related Topics 递归 记忆化搜索 数学 动态规划
// 👍 289 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//f(n) = f(n-1) + f(n-2)
func fib(n int) int {
	return dfs(n)
}

//动态规划
func fibN(n int) int {
	if n < 2 {
		return n
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 1

	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}

//递归实现
//缓存
var m map[int]int = make(map[int]int)

func dfs(n int) int {
	if n < 2 {
		return n
	}

	if _, ok := m[n]; ok {
		return m[n]
	}

	m[n] = dfs(n-1) + dfs(n-2)
	return m[n]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(fib(3))
}
