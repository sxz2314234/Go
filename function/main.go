/*

1. value transfer
2. qution tranfer
3. function assignment

*/

package main

import (
	"fmt"
	// "func/src/func"
	"func/src/array"
	// "math"
)

func main() {
	// getSquareroot := func(square float64) float64 { return math.Sqrt(float64(square)) }
	// var value1, value2 int = 2, 3

	// fmt.Printf("Before swaping...\n")
	// print(value1, value2)
	// fmt.Printf("\n")

	// value1, value2 = _func.Swap(value1, value2)
	// fmt.Printf("After swaping...\n")
	// _func.Print(value1, value2)
	// fmt.Printf("\n")

	// _func.Swap1(&value1, &value2)
	// fmt.Printf("After swaping1...\n")
	// _func.Print(value1, value2)
	// fmt.Printf("\n")

	// var root float64 = 9
	// fmt.Printf("Now, the root is %f\n", root)
	// fmt.Printf("After squaring: %f", getSquareroot(root))

	// nextNum := _func.Getsequence()
	// fmt.Println(nextNum())
	// fmt.Println(nextNum())
	// fmt.Println(nextNum())

	// nextNum = _func.Getsequence()
	// fmt.Println(nextNum())
	// fmt.Println(nextNum())
	// fmt.Println(nextNum())

	// 方法
	// var c _func.Circle
	// c.Radius = 10.0
	// fmt.Println("Area of Circle(c1) = ", c.GetArea())

	var arr1 = [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[:]
	fmt.Println(array.Sum(slice1, 5))
}
