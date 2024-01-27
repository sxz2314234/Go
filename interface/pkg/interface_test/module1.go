// polymorphism

package interface_test

import "fmt"

type Animals interface {
	Groan()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Groan() {
	fmt.Println("The dog's shout is wangwangwang~~~")
}
func (c Cat) Groan() {
	fmt.Println("The cat's shout is miaomiaomiao~~~")
}
