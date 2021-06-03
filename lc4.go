package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMiddleNum(nums1, nums2)
}
//leetcode submit region end(Prohibit modification and deletion)
//log(m+n)
func findMiddleNum(nums1, nums2 []int) float64 {
	//1 3 6 9
	//2 4 5 8 10
	l1, l2 := len(nums1), len(nums2)


	//mid = (l1+l2)/2
	//odd  num = nums[mid]
	//even num = （nums[mid] + nums[mid]）/ 2

	//先考虑长度为奇数的情况，即在两个有序数组中找到第mid+1个元素

	mid := (l1+l2)/2
	if (l1+l2)%2 == 0 {
		sum := findKthNum(nums1, nums2, 0, 0, mid) + findKthNum(nums1, nums2, 0, 0, mid+1)
		return float64(sum)/2.0
	} else {
		return float64(findKthNum(nums1, nums2, 0, 0, mid+1))
	}
}

func findKthNum(nums1, nums2 []int, i, j int, k int) int {
	l1, l2 := len(nums1), len(nums2)

	if i == l1 {
		return nums2[j+k-1]
	}
	if j == l2 {
		return nums1[i+k-1]
	}

	if k == 1 {
		if i < l1 && j < l2 {
			return min2num(nums1[i], nums2[j])
		} else if i < l1 {
			return nums1[i]
		} else if j < l2 {
			return nums2[j]
		} else {
			return -1
		}
	}

	//nums exclude pre (k/2) num
	half_k := k/2
	idx1 := min2num(half_k+i-1, l1-1)
	n1 := nums1[idx1]

	idx2 := min2num(half_k+j-1, l2-1)
	n2 := nums2[idx2]

	if n1 < n2 {
		//exclude nums1 pre (k/2) num
		return findKthNum(nums1, nums2, idx1+1, j, k - (idx1-i+1))
	} else {
		//exclude nums2 pre k/2 num
		return findKthNum(nums1, nums2, i, idx2+1, k - (idx2-j+1))
	}
}

func min2num(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(findMiddleNum([]int{1}, []int{2, 3, 4, 5, 6}))
}