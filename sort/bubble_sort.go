package main

import "fmt"

//冒泡排序是一种简单的排序算法。它重复访问要排序的数组，一次比较两个元素，如果它们的顺序错误就把他们交换过来。
//重复直到待排序数组为空为止，也就是整个数组有序。
//这个算法的名字由来是因为越小的元素会经过交换慢慢的"浮"到数组的顶端
func BubbleSort(nums []int) []int {
	//重复n次查找，每次查找都能找到一个最大的值，放到末尾
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

func main() {
	fmt.Println(BubbleSort([]int{3, 2, 5, 1}))
}
