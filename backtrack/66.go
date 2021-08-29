package main

import "fmt"

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//你可以假设除了整数 0 之外，这个整数不会以零开头。
func plusOne(digits []int) []int {
	var check func(nums []int, carry int)
	check = func(nums []int, carry int) {
		if carry == 0 {
			return
		}
		l := len(nums)
		//carry > 0
		if l == 0 {
			nums = append([]int{carry}, nums...)
			return
		}

		sum := nums[l-1] + carry
		if sum < 10 {
			nums[l-1] = sum
			return
		} else {
			nums[l-1] = sum % 10
			carry = sum / 10
			check(nums[:l-1], carry)
			nums = append(nums[:l-1], nums[l-1])
			return
		}
	}

	check(digits, 1)
	// if carry > 0 {
	//     digits = append([]int{carry}, digits...)
	// }
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}
