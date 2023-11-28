package datastructures

import (
	"fmt"
	"time"
)

func DemonstrateChannels() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Message from channel!"
	}()

	fmt.Println("Channel demonstration:")
	msg := <-ch // Receive message from channel
	fmt.Println(msg)
}
