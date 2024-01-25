// 按照要求求出列或者行的代数和
// 1: 求行; 2: 求列

package multiarray

import (
	"fmt"
)

func Multi_sum(arr [][]int, n int, sign int) []int {
	var result = make([]int, n)
	switch sign {
	case 1:
		{
			for i := 0; i < n; i++ {
				result[i] = 0
				for j := 0; j < n; j++ {
					result[i] += arr[i][j]
				}
			}
		}
	case 2:
		{
			for j := 0; j < n; j++ {
				result[j] = 0
				for i := 0; i < n; i++ {
					result[j] += arr[i][j]
				}
			}
		}
	default:
		fmt.Print("Please enter 1 or 2\n")
	}
	return result
}

func Bianli(arr [][]int) {
	for i, x := range arr {
		fmt.Printf("The %dth is: %d\n", i, x)
	}
}
