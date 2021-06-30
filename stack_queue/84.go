package main

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
// 示例:
//
// 输入: [2,1,5,6,2,3]
//输出: 10
// Related Topics 栈 数组 单调栈
// 👍 1413 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//思路：单调栈的应用，求当前柱子前面的面积，构建一个单调递增栈
//从左往右找到第一个递减节点为当前节点，因为递减，所以当前节点无法按照前一个节点的高度来计算面积，
//这时，我们就可以计算当前节点前面那一部分递增区间的面积
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	ret := 0
	//记录当前值左边的节点下标
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}

		//ret = max2num(ret, cur*1)
		for len(stack) != 0 && cur < heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			h := heights[pop]
			//计算当前节点前面部分的宽度
			var w int
			//如果栈里没有元素，pop元素为高度最小的元素，那么宽度即等于i
			if len(stack) != 0 {
				//从peek到cur节点都可以按照h的高度来计算面积
				peek := stack[len(stack)-1]
				w = i - peek - 1
			} else {
				w = i
			}
			ret = max2num(ret, h*w)
		}

		stack = append(stack, i)
	}

	return ret
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(largestRectangleArea(height))
}
