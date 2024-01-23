// 闭包, 返回函数

package _func

func Getsequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 方法, 指定接收对象
type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
}
