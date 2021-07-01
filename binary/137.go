package main

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,3,2]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,1,0,1,99]
//输出：99
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
// Related Topics 位运算 数组
// 👍 685 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumberV2(nums []int) int {
	ret := 0
	//依次统计所有元素的每一位1的个数，对3取余得到的数（0/1）就是结果的那一位的值
	for i := 0; i < 64; i++ {
		sum := 0
		for _, v := range nums {
			sum += (v >> i) & 1
		}
		ret = ret ^ (sum%3)<<i
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
