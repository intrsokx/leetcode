package main

import "fmt"

//思路：把一个数组分为左右两部分，左边的元素小于右边的元素
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	//选取最后一个元素作为基准元素
	p := nums[right]

	i := left
	for j := left; j < right; j++ {
		if nums[j] < p {
			//swap
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	//将基准元素赋换到中间
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {
	//思想：将数组分成左右两部分，左半部分小于右半部分
	fmt.Println(QuickSort([]int{2, 3, 1, 2, 5, 6, 7, 3}))
}
