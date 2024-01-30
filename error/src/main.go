package main

import (
	"error/pkg/error_mechanism"
	"fmt"
)

func main() {
	// correct case
	if result, errorMsg := errormechanism.Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	// error case
	if _, errorMsg := errormechanism.Divide(100, 0); errorMsg != "" {
		fmt.Println("The errorMsg is ", errorMsg)
	}
}
