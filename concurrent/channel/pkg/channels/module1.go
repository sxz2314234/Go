/*

declare: channel type
var varname chanel vartype

*/

package typechannel

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

type Channel interface {
	Receiver(int)
	Sender(chan int)
}

type Chaner struct {
	Size     int
	Channels chan int
}

func (ch *Chaner) Receiver(id int) {
	defer Wg.Done()

	value := <-ch.Channels
	fmt.Printf("Receiver %d received: %d\n", id, value)
	// ch.Size--
	// fmt.Printf("There are still %d busy space\n", ch.Size)
}

func (ch *Chaner) Sender(value *int) {
	defer Wg.Done()

	ch.Channels <- *value
	// ch.Size++
	// fmt.Printf("The buffer has %d vacant space\n", 4-ch.Size)
}
