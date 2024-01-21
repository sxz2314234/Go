package main

import "fmt"

var a string = "Hello World"
var b bool = true

func main() {
	var x int = 3
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型是 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int)")
	case bool, string:
		fmt.Printf("x 是 bool或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
