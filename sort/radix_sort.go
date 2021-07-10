package main

import "fmt"

//基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。
//有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。
//基数排序：基数排序是按照低位先排序，然后收集（按当前位顺序重新写入数组），然后再高位排序，然后再收集；依次类推，直到最高位
func RadixSort(nums []int) []int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	maxDigit := getMaxDigit(max)

	mod := 10
	dev := 1
	for i := 0; i < maxDigit; i, mod, dev = i+1, mod*10, dev*10 {
		counter := make([][]int, 10)
		for _, v := range nums {
			bucket := (v % mod) / dev
			counter[bucket] = append(counter[bucket], v)
		}

		idx := 0
		for j := 0; j < len(counter); j++ {
			for _, v := range counter[j] {
				nums[idx] = v
				idx++
			}
		}
	}

	return nums
}

func getMaxDigit(max int) int {
	c := 0
	for max != 0 {
		max = max / 10
		c++
	}
	return c
}

func main() {
	fmt.Println(RadixSort([]int{2, 12, 23, 33, 4, 5, 45, 5, 6}))
}
