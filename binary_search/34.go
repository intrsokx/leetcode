package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(findLeftIdx(nums, 1))
	fmt.Println(findRightIdx(nums, 1))

	fmt.Println(searchInsertV2([]int{1, 1, 1}, 1))
	fmt.Println(1/1, 1%1)

	fmt.Println(searchV3([]int{3, 1, 1}, 1))
}

// 找到第一个大于等于target的位置，就是该元素的插入位置
func searchInsertV2(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := (end-start)/2 + start
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			//end向左靠，就是找到最左边的插入位置
			//如果是start向右靠，就是找到最右边的插入位置
			end = mid
		}
	}
	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else {
		//target 比end还大
		return end + 1
	}
}

func findLeftIdx(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	//start+1 = end
	for start+1 < end {
		mid := (end-start)/2 + start

		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			//move end to left
			end = mid
		}
	}

	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	} else {
		return -1
	}
}

func findRightIdx(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1

	//[start, end] >= 3
	for start+1 < end {
		mid := (end-start)/2 + start
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			//move start to right
			start = mid
		}
	}

	if nums[end] == target {
		return end
	} else if nums[start] == target {
		return start
	} else {
		return -1
	}
}

func searchV3(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	start, end := 0, len(nums)-1
	for start+1 < end {
		//deal repeat item
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		for start < end && nums[end] == nums[end-1] {
			end--
		}

		mid := (end-start)/2 + start
		if nums[mid] == target {
			return true
		}
		//两段上升的数组，判断mid、target在哪部分
		if nums[mid] >= nums[start] {
			//[start, mid]
			if target >= nums[start] && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[mid] <= nums[end] {
			//[mid, end]
			if target >= nums[mid] && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}
