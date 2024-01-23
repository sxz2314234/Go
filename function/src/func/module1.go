package _func

import (
	"fmt"
)

func Swap(i int, j int) (int, int) { return j, i }

func Swap1(i *int, j *int) {
	var temp int
	temp = *i
	*i = *j
	*j = temp
}

func Print(i int, j int) {
	fmt.Printf("The first number is: %d\n", i)
	fmt.Printf("The second number is: %d\n", j)
}
