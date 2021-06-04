package benching

import (
	binary_heap "github.com/samarec1812/golang-data-structure/binary-heap"
	"math/rand"
	"testing"
)
func GetMaxBench(n int, b *testing.B) {
	heap := binary_heap.BinaryHeap{}
	values := make([]int, 0, 0)
	for i := 0; i < 10000; i++ {
		values = append(values, rand.Int()%150000)
	}
	heap.BuildHeap(values)
	for i := 0; i < n; i++ {
		heap.GetMax()
	}
}

func BenchmarkHeapGetMax10x1(b *testing.B) {
	//first := time.Now()
	for i := 0; i < b.N; i++ {
		GetMaxBench(10, b)
	}
	//second := time.Now()
	//fmt.Print(second.Sub(first))
}

func BenchmarkHeapGetMax10x2(b *testing.B) {
	//first := time.Now()
	for i := 0; i < b.N; i++ {
		GetMaxBench(100, b)
	}
	//second := time.Now()
	//fmt.Print(second.Sub(first))
}

func BenchmarkHeapGetMax10x3(b *testing.B) {
	//first := time.Now()
	for i := 0; i < b.N; i++ {
		GetMaxBench(1000, b)
	}
	//second := time.Now()
	//fmt.Print(second.Sub(first))
}

func BenchmarkHeapGetMax10x4(b *testing.B) {
	//first := time.Now()
	for i := 0; i < b.N; i++ {
		GetMaxBench(10000, b)
	}
	//second := time.Now()
	//fmt.Print(second.Sub(first))
}

