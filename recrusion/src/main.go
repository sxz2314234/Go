package main

import (
	"fmt"
	"recrusion/pkg/recrusion"
)

func main() {
	var test = 4
	fmt.Println(recrusion.Factorial(test))
	fmt.Println(recrusion.Fibonacci(test))
}
