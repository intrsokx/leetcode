package main

import "fmt"

//简介：1959年由shell发明，是第一个突破O(n^2)的排序算法；是简单插入排序的改进版，它优先比较距离较远的两个元素；希尔排序又叫缩小增量排序。
//思路：先将整个序列分成若干个子序列，然后分别进行插入排序，
//希尔排序的效率提升之处在于通过对子序列的排序，使得整个序列基本有序，这样就减少了插入时元素移动的次数，从而提升了效率
func ShellSort(nums []int) []int {
	k := len(nums) / 2

	for k > 0 {
		for i := k; i < len(nums); i = i + k {
			pre := i - k
			cur := nums[i]
			//move
			for pre >= 0 && nums[pre] > cur {
				nums[pre+k] = nums[pre]
				pre = pre - k
			}
			nums[pre+k] = cur
		}
		k = k / 2
	}

	return nums
}

func main() {
	fmt.Println(ShellSort([]int{3, 2, 5, 1, 4, 3}))
}
