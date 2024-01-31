package main

import (
	"channel_test/pkg/channels"
	"fmt"
	// "fmt"
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

func main() {
	var ch = typechannel.Chaner{
		Channels: make(chan int, 2),
		Size:     0,
	}

	for i := 0; i < 4; i++ {
		typechannel.Wg.Add(1)
		go ch.Receiver(i)
	}

	for i := 0; i < 4; i++ {
		typechannel.Wg.Add(1)
		value := i + 1
		go ch.Sender(&value)
		fmt.Printf("Sent: %d\n", value)
	}

	typechannel.Wg.Wait()
	close(ch.Channels)
}
