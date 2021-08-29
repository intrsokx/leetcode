package main

import "fmt"

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)

	var check func(cur int)
	check = func(cur int) {
		if cur > n+1 {
			return
		}

		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ret = append(ret, tmp)
			return
		}

		//如果当前长度 加上 可以选择的长度小于k，说明不能构成长度为k的结果(剪枝)
		if len(path)+(n-cur+1) < k {
			return
		}

		//对当前值 选 与 不选
		path = append(path, cur)
		check(cur + 1)

		path = path[:len(path)-1]
		check(cur + 1)
	}
	check(1)
	return ret
}

func main() {
	fmt.Println(combine(4, 2))
}
