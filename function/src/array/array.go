package array

// 向函数传递数组
func Sum(array []int, n int) int {
	var result int

	for i := 0; i < n; i++ {
		result += array[i]
	}

	return result
}
