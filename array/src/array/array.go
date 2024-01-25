package array

import (
	"fmt"
)

// 向函数传递数组
func Sum(array []int, n int) int {
	var result int

	for i := 0; i < n; i++ {
		result += array[i]
	}

	return result
}

// for循环遍历
func Bianli(arr []int) {
	for i, x := range arr {
		fmt.Printf("The %dth is: %d\n", i, x)
	}
}
