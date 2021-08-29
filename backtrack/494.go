package main

import "fmt"

//给你一个整数数组 nums 和一个整数 target 。
//向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//求这个表达式等于target的个数

//对于每个元素都有两种情况（+-），那么递归进入函数，求sum，当所有元素都被求完后，判断sum是否等于target，若等于，则全局的计数器加加
func findTargetSumWays(nums []int, target int) int {
	var cnt int
	var check func(i, sum int)
	check = func(i, sum int) {
		if i == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}

		check(i+1, sum+nums[i])
		check(i+1, sum-nums[i])
	}

	check(0, 0)
	return cnt
}

func main() {
	fmt.Println(findTargetSumWays([]int{1}, 1))
}
