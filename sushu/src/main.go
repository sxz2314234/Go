package main

import "fmt"

func sushu(obj int) {
	var i, j int

	for i = 2; i < obj; i++ {
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}

		if j == i {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}

func main() {
	var count int = 100

	sushu(count)
}
