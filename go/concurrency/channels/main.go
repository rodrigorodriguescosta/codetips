package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channelExample()
}

func readChat(ch <-chan string) {
	ok := true
	var chValue string
	for ok {
		chValue, ok = <-ch
		fmt.Println("Received value: ", chValue)
	}
}

func sendMessage(ch chan<- string, chState chan<- bool) {
	times := 0
	messages := []string{"Hello", "Hi", "Are you sure?", "Go is awesome"}
	for times < 10 {
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		ch <- messages[rand.Intn(3)]
		times++
	}
	close(ch)
	chState <- true
}

func channelExample() {
	chChat := make(chan string)
	chChatIsClosed := make(chan bool)
	go readChat(chChat)
	go sendMessage(chChat, chChatIsClosed)
	<-chChatIsClosed
	fmt.Println("Channel example is Done")
}
