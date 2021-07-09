package main

import "fmt"

//计数排序：计数排序不是基于比较的排序算法， 其核心思想在于将数据转换为键值存储在额外的空间中；这个空间可以是数组或者map。
//因为需要预先申请确定的空间，因此，计数排序要求输入的数据必须是有确定范围内的整数。
func CountSort(nums []int, max, min int) []int {
	bucket := make(map[int]int, max-min+1)

	for i := min; i <= max; i++ {
		bucket[i] = 0
	}

	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}

	idx := 0
	for i := min; i <= max; i++ {
		if bucket[i] > 0 {
			for j := 0; j < bucket[i]; j++ {
				nums[idx] = i
				idx++
			}
		}
	}

	return nums
}

//数组元素值范围[0，max]
func CountSortV2(nums []int, max int) []int {
	bucket := make([]int, max+1)

	for _, num := range nums {
		bucket[num]++
	}

	idx := 0
	for i := 0; i < len(bucket); i++ {
		if bucket[i] > 0 {
			for j := 0; j < bucket[i]; j++ {
				nums[idx] = i
				idx++
			}
		}
	}

	return nums
}

func main() {
	fmt.Println(CountSort([]int{3, 2, 3, 1, 0, -1, -2}, 3, -2))
	fmt.Println(CountSortV2([]int{3, 2, 1, 3, 2}, 3))
}
