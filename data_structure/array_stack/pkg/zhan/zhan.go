package zhan

type Student struct {
	Name string
	Sno  int
}

type Stack struct {
	arr []*Student
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func Newstack() *Stack {
	return &Stack{
		arr: make([]*Student, 0, 16),
	}
}

func (s *Stack) Push(stu *Student) {
	// slice can entend automatically
	s.arr = append(s.arr, stu)
}

func (s *Stack) Top() *Student {
	if s.Empty() {
		return nil
	}

	return s.arr[len(s.arr)-1]

}

func (s *Stack) Pop() *Student {
	if s.Empty() {
		return nil
	}

	val := s.Top()
	s.arr = s.arr[:len(s.arr)-1]
	return val
}
