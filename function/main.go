/*

1. value transfer
2. qution tranfer
3. function assignment

*/

package main

import (
	"fmt"
	"func/src/_func"
	"math"
)

func main() {
	getSquareroot := func(square float64) float64 { return math.Sqrt(float64(square)) }
	var value1, value2 int = 2, 3

	fmt.Printf("Before swaping...\n")
	print(value1, value2)
	fmt.Printf("\n")

	value1, value2 = _func.Swap(value1, value2)
	fmt.Printf("After swaping...\n")
	_func.Print(value1, value2)
	fmt.Printf("\n")

	_func.Swap1(&value1, &value2)
	fmt.Printf("After swaping1...\n")
	_func.Print(value1, value2)
	fmt.Printf("\n")

	var root float64 = 9
	fmt.Printf("Now, the root is %f\n", root)
	fmt.Printf("After squaring: %f", getSquareroot(root))

}
