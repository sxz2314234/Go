package main

import (
	"channel_test/pkg/channels"
	"fmt"
	"sync"
)

/*
	func receive(channel chan int) {
		rec := <-channel
		fmt.Println("接受成功! ", rec)
	}

	func main() {
		// channel with no buffer
		channel := make(chan int)
		go receive(channel)
		channel <- 10
		fmt.Println("发送成功!")

}
*/

var wg sync.WaitGroup

func main() {
	var ch=typechannel.Chaner{
		Channels: make(chan int,2),
		Size: 0,
	}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go ch.Receiver(i)
		wg.Done()

		temp:=ch.Sender()
		fmt.Println() go ch.Sender()
	}
}
