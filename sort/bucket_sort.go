package main

import "fmt"

//桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。
//桶排序 (Bucket sort)的工作的原理：假设输入数据服从均匀分布，将数据分到有限数量的桶里，
//每个桶再分别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排）。

//思路：
//1、设置一个定量的数组当做空桶； 2、遍历输入数据，并把数据一个一个放到对应的桶里去；
//3、对每个不是空的桶进行排序； 4、按照顺序把所有桶的数据拼接起来

//算法分析：桶排序最好情况下是线性时间O(n)，桶排序的时间复杂度，取决于各个桶之间数据进行排序的时间复杂度，因为其它部分的时间复杂度都为O(n)。
//所以，桶划分的越小，也就是各个桶之间的数据越少，排序所用的时间就会越少，但同时空间消耗就会增大。

//bucketSize:每个桶的大小
//bucketCount:有多少个桶
//min:待排序列中的最小值
//max:待排序列中的最大值
func BucketSort(nums []int, bucketSize int) []int {
	if len(nums) == 0 {
		return nums
	}

	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	defaultBucketSize := 5

	//最小为5
	bucketSize = defaultBucketSize | bucketSize

	bucketCount := (max-min)/bucketSize + 1

	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		//因为下面是append操作，所以这里初始化数组容量是0
		buckets[i] = make([]int, 0)
	}

	//根据映射函数将数据分配到桶中，桶0中的元素全部小于桶1中的元素
	for _, v := range nums {
		//v-min 是为了预防出现负数的情况
		idx := (v - min) / bucketSize
		buckets[idx] = append(buckets[idx], v)
	}

	//最后再分别对bucket中的各个数组进行排序，然后合并
	ret := make([]int, 0)
	for i := 0; i < len(buckets); i++ {
		insertionSort(buckets[i])

		ret = append(ret, buckets[i]...)
	}

	return ret
}

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		pre := i - 1
		cur := nums[i]
		for pre >= 0 && cur < nums[pre] {
			nums[pre+1] = nums[pre]
			pre--
		}
		nums[pre+1] = cur
	}

	return nums
}

func main() {
	fmt.Println(BucketSort([]int{3, 4, 5, -1, 2, 7, 6, 8, 9}, 5))
}
