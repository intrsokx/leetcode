package main

import "fmt"

//思路：将数组分为有序无序两部分，对于无序部分数据，从后向前扫描有序部分，寻找相应的位置插入
func InsertionSort(nums []int) []int {
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
	fmt.Println(InsertionSort([]int{3, 2, 4, 0, 1}))
}
