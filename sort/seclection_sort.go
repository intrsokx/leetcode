package main

import "fmt"

//思路：将数组分为有序无序两部分，有序部分初始为空，每次从无序部分中找出最小的元素，放到有序部分的末尾；重复这个过程，直到无序部分为空为止
func SelectionSort(nums []int) []int {
	var min int
	for i := 0; i < len(nums); i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}

		//swap
		nums[min], nums[i] = nums[i], nums[min]
	}

	return nums
}

func main() {
	fmt.Println(SelectionSort([]int{3, 2, 4, 1, 5}))
}
