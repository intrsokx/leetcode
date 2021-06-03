package main

import "fmt"

/**
分治法应用
*/
// V2：通过分治法遍历二叉树
func preorderTraversal(root *TreeNode) []interface{} {
	result := divideAndConquer(root)
	return result
}
func divideAndConquer(node *TreeNode) []interface{} {
	ret := make([]interface{}, 0)
	if node == nil {
		return ret
	}

	//divide
	left := divideAndConquer(node.Left)
	right := divideAndConquer(node.Right)

	//conquer
	ret	= append(ret, node.Val)
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret
}

//归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) >> 1
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}
//left && right 分别是有序的
func merge(left, right []int) []int {
	ret := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			ret = append(ret, left[l])
			l++
		} else {
			ret = append(ret, right[r])
			r++
		}
	}

	//因为nums[len(nums):]不会报错，而是返回一个长度为0的切片，所以这块可以不用判断l和r的长度，直接对ret追加
	ret	= append(ret, left[l:]...)
	ret = append(ret, right[r:]...)

	return ret
}



/**
快排由于是原地交换所以没有合并过程 传入的索引是存在的索引（如：0、length-1 等），如果越界的话，会导致程序崩溃
*/
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, 0, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	//比较的基准值
	p := nums[start]
	idx := start

	for i := start+1; i <= end; i++ {
		if nums[i] < p {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	//把中间的值替换为基准值
	nums[idx] = p
	return idx
}

//go run tree.go divide.go
func main() {
	nums := []int{1, 4, 5, 2, 3}

	fmt.Println(MergeSort(nums))
	fmt.Println(QuickSort(nums))
}
