package main

import "fmt"

//思路：堆排序的核心就是将数组构建一个最大堆（完全二叉树），然后每次将堆顶元素取出放到数组的末尾；
//重复以上过程，直到待排序数组长度为1为止
func HeapSort(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		buildSink(nums, i, len(nums))
	}

	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]

		buildSink(nums, 0, i)
	}

	return nums
}

func buildSink(nums []int, i, length int) {
	for {
		l := 2*i + 1
		r := 2*i + 2

		idx := i
		if l < length && nums[l] > nums[idx] {
			idx = l
		}
		if r < length && nums[r] > nums[idx] {
			idx = r
		}

		//当前节点值比子节点大，不用下沉
		if idx == i {
			break
		}

		//swap && sink
		nums[idx], nums[i] = nums[i], nums[idx]
		i = idx
	}
}

func main() {
	fmt.Println(HeapSort([]int{4, 10, 3, 5, 1}))
}
