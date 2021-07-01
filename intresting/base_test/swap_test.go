package base_test

import (
	"testing"
)

func swapByXor(a, b int) (int, int) {
	//a, b = b, a
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func swap_v2(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func swap_v1(a, b int) (int, int) {
	var temp int
	a = temp
	a = b
	b = temp
	return a, b
}

func BenchmarkSwapWithXor(b *testing.B) {
	x, y := 200, 8888
	for i := 0; i < b.N; i++ {
		x = x ^ y
		y = x ^ y
		x = x ^ y
	}
}

func BenchmarkSwapV1(b *testing.B) {
	x, y, tmp := 200, 8888, 0
	for i := 0; i < b.N; i++ {
		tmp = x
		x = y
		y = tmp
	}
}

func BenchmarkSwapV2(b *testing.B) {
	x, y := 200, 8888
	for i := 0; i < b.N; i++ {
		x, y = y, x
	}
}
