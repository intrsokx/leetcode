package main

import "fmt"

func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	//divide
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	result := merge(left, right)
	return result
}

func merge(left, right []int) []int {
	ret := make([]int, len(left)+len(right))

	i, j := 0, 0
	idx := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret[idx] = left[i]
			idx++
			i++
		} else {
			ret[idx] = right[j]
			idx++
			j++
		}
	}

	copy(ret[idx:], left[i:])
	copy(ret[idx:], right[j:])

	return ret
}

func main() {
	ret := []int{1, 2, 3, 0, 0, 0}
	other := []int{4, 5, 6}
	//ret = append(ret, other...)
	copy(ret[3:], other[3:])
	copy(ret[3:], other)
	fmt.Println(ret)

	fmt.Println(MergeSort([]int{2, 3, 1, 0, 5, 7, 6}))
}
