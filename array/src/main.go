package main

import (
	"array/pkg/array"
	"array/pkg/multi_array"
	// "fmt"
)

func main() {
	var arr = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	var result []int = multiarray.Multi_sum(arr, 3, 1)

	// var arr1 = [3]int{1, 2, 3}
	// fmt.Print(array.Sum(arr1[:], 3))
	array.Bianli(result)
}
