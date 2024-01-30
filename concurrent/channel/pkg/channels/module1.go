/*

declare: channel type
var varname chanel vartype

*/

package typechannel

import (
	"fmt"
)

type Channel interface {
	Receiver(int)
	Sender() int
}

type Chaner struct {
	Size     int
	Channels chan int
}

func (ch *Chaner) Receiver(i int) {
	fmt.Println("成功接收!!!")
	ch.Channels <- i
	ch.Size++
	fmt.Println("现在的通道数为: ", ch.Size)
}

func (ch *Chaner) Sender() int {
	fmt.Println("成功发送!!!")
	ch.Size--
	fmt.Println("现在的通道数: ", ch.Size)
	return <-ch.Channels
}
