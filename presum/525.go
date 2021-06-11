package main

import "fmt"

//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
//
//
//
// 示例 1:
//
//
//输入: nums = [0,1]
//输出: 2
//说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
//
// 示例 2:
//
//
//输入: nums = [0,1,0]
//输出: 2
//说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
//
// Related Topics 哈希表
// 👍 407 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
	分析：连续子数组中1和0的个数是相同的，可推理出nums[i] + ... + nums[j] == (j-i+1)/2
假设我们把0当成-1，即变成了求各个元素之和为0的最长连续子数组。
*/
func findMaxLength0(nums []int) int {
	//        sum idx
	//mp := map[int]int{-1: 0}
	//max := 0
	//sum := 0
	//for i, n := range nums {
	//	if n == 1 {
	//		sum++
	//	} else {
	//		sum--
	//	}
	//
	//	mp[sum] = i
	//}

	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	//[i, j]  =>  preSum[j] - preSum[i] + nums[i] == 0
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			l := j - i + 1
			sum := preSum[j] - preSum[i] + nums[i]
			if l%2 == 0 && sum == l/2 {
				max = max2num(max, l)
			}
		}
	}

	return max
}

func max2num(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	//map[sum]idx
	mp := map[int]int{0: -1}
	max := 0
	sum := 0
	for i, n := range nums {
		if n == 1 {
			sum++
		} else {
			sum--
		}

		if idx, ok := mp[sum]; ok {
			//说明preSum[idx] == preSum[i] && idx < i
			//可推导出 preSum[i] - preSum[idx] == 0 ==> nums[idx+1, j] 是和为0的子数组
			max = max2num(max, i-idx-1+1)
		} else {
			mp[sum] = i
		}
	}

	return max
}

var version string

//go build -ldflags '-X main.version="dev"'
//go run -ldflags '-X main.version="dev"' 525.go
func main() {
	fmt.Println("version:", version)
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}

//leetcode submit region end(Prohibit modification and deletion)
