/*

1. value transfer
2. qution tranfer
3. function assignment

*/

package main

import (
	"fmt"
	"math"
)

func swap(i int, j int) (int, int) { return j, i }

func swap1(i *int, j *int) {
	var temp int
	temp = *i
	*i = *j
	*j = temp
}

func print(i int, j int) {
	fmt.Printf("The first number is: %d\n", i)
	fmt.Printf("The second number is: %d\n", j)
}

func main() {
	getSquareroot := func(square float64) float64 { return math.Sqrt(float64(square)) }
	var value1, value2 int = 2, 3

	fmt.Printf("Before swaping...\n")
	print(value1, value2)
	fmt.Printf("\n")

	value1, value2 = swap(value1, value2)
	fmt.Printf("After swaping...\n")
	print(value1, value2)
	fmt.Printf("\n")

	swap1(&value1, &value2)
	fmt.Printf("After swaping1...\n")
	print(value1, value2)
	fmt.Printf("\n")

	var root float64 = 9
	fmt.Printf("Now, the root is %f\n", root)
	fmt.Printf("After squaring: %f", getSquareroot(root))
}
