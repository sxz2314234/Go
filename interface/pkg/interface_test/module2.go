// decoupling and: 解耦和

package interface_test

import "fmt"

type Printer interface {
	Print()
}

type ConsolePrinter struct {
	Name string
	Info int
}

type FilePrinter struct {
	Name string
	Info int
}

func (cp ConsolePrinter) Print() {
	fmt.Println("This is ", cp.Name)
	fmt.Println("The information is ", cp.Info)
}

func (fp FilePrinter) Print() {
	fmt.Println("This is ", fp.Name)
	fmt.Println("The information is ", fp.Info)
}
