package benchmark

import (
	"testing"
	"time"
)

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}

func BenchmarkSliceNoCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceNoCap()
	}
}

func BenchmarkSliceWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceWithCap()
	}
}

func BenchmarkFib(b *testing.B) {
	time.Sleep(3 * time.Second) // 模拟耗时的准备工作
	b.ResetTimer()              // 重置计时器，忽略前面的准备时间
	for n := 0; n < b.N; n++ {
		sum(10, 1)
	}
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()               // 暂停计时
		time.Sleep(1 * time.Second) // 每次函数执行前的准备工作
		b.StartTimer()              // 继续计时
		sum(10, 1)                  // 被测函数
	}
}

//D:\goitem\go-box\practice\benchmark>go test -bench=.
//goos: windows
//goarch: amd64
//pkg: benchmark
//cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//BenchmarkSum-12         1000000000               0.2382 ns/op
//PASS
//ok      benchmark       0.734s

//$ go test -bench=. -count=3
//goos: windows
//goarch: amd64
//pkg: benchmark
//cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//BenchmarkSum-12         1000000000               0.2435 ns/op
//BenchmarkSum-12         1000000000               0.2392 ns/op
//BenchmarkSum-12         1000000000               0.2534 ns/op
//PASS
//ok      benchmark       1.122s

//$  go test -bench=. -cpu=1,4,6,10,12
//goos: windows
//goarch: amd64
//pkg: benchmark
//cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//BenchmarkSum            1000000000               0.2393 ns/op
//BenchmarkSum-4          1000000000               0.2399 ns/op
//BenchmarkSum-6          1000000000               0.2443 ns/op
//BenchmarkSum-10         1000000000               0.2424 ns/op
//BenchmarkSum-12         1000000000               0.2473 ns/op
//PASS
//ok      benchmark       1.656s

//$  go test -bench='Cap$' -benchmem .
//goos: windows
//goarch: amd64
//pkg: benchmark
//cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//BenchmarkSliceNoCap-12             17128             68519 ns/op          386297 B/op         20 allocs/op
//BenchmarkSliceWithCap-12           69754             16739 ns/op           81920 B/op          1 allocs/op
//PASS
//ok      benchmark       3.730s

//$ go test -bench='Fib$'
//goos: windows
//goarch: amd64
//pkg: benchmark
//cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//BenchmarkFib-12         1000000000               0.2373 ns/op
//PASS
//ok      benchmark       18.570s
