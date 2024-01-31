package main

import (
	// "channel_test/pkg/channels"
	"channel_test/pkg/mutex_practise"
	"fmt"
	"time"
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
/*
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

	ch2 := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch2:
			fmt.Println(x)
		case ch2 <- i:
		}
	}
}
*/

func main() {
	start := time.Now()
	for i := 0; i < 11; i++ {
		go mutexpractise.Write()
		mutexpractise.Wg2.Add(1)
	}

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		go mutexpractise.Read()
		mutexpractise.Wg2.Add(1)
	}

	mutexpractise.Wg2.Wait()
	fmt.Println(time.Since(start))
}
