package base_test

import "testing"

//计算斐波那契数列中的第n个数字
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

//go test -bench BenchmarkFib20 -cpu=1,2,4 -benchtime=10s
//go test -bench BenchmarkFib20 -count=10 |tee fib.txt

//benchstat fib.txt

/**
benchstat 告诉我们，平均值为52.4微秒，样本间的波动区间为正负 2％。 这对电池电量来说在意料之中。

第一次运行是最慢的，因为操作系统的 CPU 时钟频率已经降低以节省功耗。
接下来的两次运行是最快的，因为操作系统识别到有一个较大的工作负载加入，就会提高 CPU 时钟速度，以尽快通过工作。
剩下的是当 CPU 高速运转发热，因为功耗导致又被限制，所以又慢了下来
*/

//优化比较：
//go test -c 得到base_test.test 重命名为bast_test.golden
//./base_test.golden -test.bench=Fib -test.count=10 > old.txt
//./base_test.test -test.bench=Fib -test.count=10 > new.txt
//比较 benchstat old.txt new.txt
//
