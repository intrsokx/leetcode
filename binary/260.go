package main

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
//
//
// 示例 2：
//
//
//输入：nums = [-1,0]
//输出：[-1,0]
//
//
// 示例 3：
//
//
//输入：nums = [0,1]
//输出：[1,0]
//
//
// 提示：
//
//
// 2 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次
//
// Related Topics 位运算 数组
// 👍 413 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//找到只出现一次的那两个元素，其余所有元素均出现两次
//思路：通过异或可以求得a = x^y，那么只要将数组分成两部分，一部分含有x,一部分含有y，分别用a异或这两部分就可得到结果
//那么怎么将数组分为这两部分呢，从二进制的角度去看，如果x y的某一位不相同，则异或后a的对应位为1，我们可以根据这个特性，
//根据a中为1的某一位来对整个数组进行分割，那么x,y必然会被分成不同的两部分，而剩下的元素中，元素值相同必然可以分到同一边。
//然后用a分别异或两部分，则可得到结果
func singleNumberV3(nums []int) []int {
	var a int
	for _, n := range nums {
		a ^= n
	}

	//找到a的最后一位1的位置
	tag := (a & (a - 1)) ^ a
	ret := []int{a, a}
	for _, n := range nums {
		if n&tag == 0 {
			ret[0] ^= n
		} else {
			ret[1] ^= n
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
