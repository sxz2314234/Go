package main

import (
	"Interface/pkg/interface_test"
)

func main() {
	// var animals = []interface_test.Animals{interface_test.Dog{}, interface_test.Cat{}}

	// for _, ani := range animals {
	// 	ani.Groan()
	// }

	var machines = []interface_test.Printer{
		interface_test.ConsolePrinter{
			Name: "cp1",
			Info: 123},
		interface_test.FilePrinter{
			Name: "fp1",
			Info: 234,
		},
	}

	for _, machine := range machines {
		machine.Print()
	}
}
