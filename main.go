package main

import (
	"fmt"
	segment_tree "github.com/samarec1812/golang-data-structure/segment-tree"
	"math/rand"
	"time"
)

func main () {

	//for j := 0; j < 100; j++ {
	//
	//	heap := binary_heap.BinaryHeap{}
	//	values := make([]int, 0, 0)
	//	for i := 0; i < 500000; i++ {
	//		values = append(values, rand.Int()%50000)
	//
	//	}
	//	start0 := time.Now()
	//	heap.BuildHeap(values)
	//	second0 := time.Now()
	//	duration0 := second0.Sub(start0).Milliseconds()
	//	fmt.Print("The time build: " + fmt.Sprint(duration0) + " ms : ")
	//
	//
	//	start := time.Now()
	//	for i := 0; i < 10000; i++ {
	//		// fmt.Println(heap.GetMax())
	//		heap.GetMax()
	//	}
	//	second := time.Now()
	//	duration := second.Sub(start).Milliseconds()
	//
	//	fmt.Print("The time getMax: " + fmt.Sprint(duration) + " ms : ")
	//
	//
	//	start2 := time.Now()
	//	for i := 0; i < 10000; i++ {
	//		// fmt.Println(heap.GetMax())
	//		heap.Push(rand.Int()%50000)
	//	}
	//	second2 := time.Now()
	//	duration2 := second2.Sub(start2).Milliseconds()
	//
	//	fmt.Println("The time Push: " + fmt.Sprint(duration2) + " ms")
	//	fmt.Println("------------------------------------------------")
	//}


	n := 100000
	for j := 0; j < 100; j++ {
		values := make([]int, 0, 0)
		for i := 0; i < n; i++ {
			values = append(values, rand.Int()%500)
		}
		valuesTree := make([]int, len(values)*4, len(values)*4)
		segmentTree := segment_tree.SegmentTree{values, valuesTree}
			start0 := time.Now()
			segmentTree.BuildTree(1, 0, n-1)
			second0 := time.Now()
			duration0 := second0.Sub(start0).Milliseconds()
			fmt.Print("The time BUILD: " + fmt.Sprint(duration0) + " ms : ")

			start := time.Now()
			for i := 0; i < 1000; i++ {

				segmentTree.GetSum(10, 70000, 1, 0, n-1)
			}
			second := time.Now()
			duration := second.Sub(start).Milliseconds()

			fmt.Println("SUM: " + fmt.Sprint(duration) + " ms")

		// fmt.Println(segmentTree.Arr)
	}
}
