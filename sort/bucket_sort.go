package main

//桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
//桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，
//每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

//bucketSize:每个桶的大小
//bucketCount:有多少个桶
//min:待排序列中的最小值
//max:待排序列中的最大值
func BucketSort(nums []int, bucketSize int) []int {
	if len(nums) == 0 {
		return nums
	}

	//min, max := nums[0], nums[0]
}
