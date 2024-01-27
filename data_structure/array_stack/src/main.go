package main

import (
	"fmt"
	"stack/pkg/zhan"
)

func print_student(stu zhan.Student) {
	fmt.Printf("The name of this student is %s, and its number is %d\n", stu.Name, stu.Sno)
}

func main() {
	var stu_list = []zhan.Student{
		{
			Name: "sxz",
			Sno:  123,
		},
		{
			Name: "sxz1",
			Sno:  234,
		},
		{
			Name: "sxz2",
			Sno:  345,
		},
	}

	var stu_zhan = zhan.Newstack()

	for i := 0; i < len(stu_list); i++ {
		print_student(stu_list[i])
		stu_zhan.Push(&stu_list[i])
	}

	fmt.Printf("\n")
	var temp_stu = make([]*zhan.Student, len(stu_list))
	for i := range stu_list {
		temp_stu[i] = stu_zhan.Pop()
	}
	for _, stu := range temp_stu {
		print_student(*stu)
	}
}
