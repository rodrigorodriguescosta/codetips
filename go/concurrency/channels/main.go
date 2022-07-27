package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channelExample()
}

// readChat function that receive a chanel(readonly)
// when we put <- vefore the channel, it means the channel is readonly
func readChat(ch <-chan string) {
	ok := true
	var chValue string
	for ok {
		// we are receiving data from the channel, it
		chValue, ok = <-ch
		fmt.Printf("%s Received value: %s \n", time.Now().Format("15:04:05"), chValue)
	}
}

func sendMessage(ch chan<- string, chState chan<- bool) {
	times := 0
	// create some random messages
	messages := []string{
		"Hello", "Hi", "Are you sure?", "Go is awesome",
		"is simply dummy text of the printing",
		"typesetting industry. Lorem Ipsum ha",
		"Contrary to popular belief",
		"There are many variations of passages of Lorem Ipsum available",
	}
	for times < 10 {
		// simulating some preocessing, it should be api call for instance
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))

		// here we are sending a string to the channel
		ch <- messages[rand.Intn(len(messages)-1)]
		times++
	}
	close(ch)
	chState <- true
}

func channelExample() {
	// create a channel(unbuffered) to transfer data bwtweeen go routine
	chChat := make(chan string)

	// this another channel(unbuffered) is used to control when the job is done
	chChatIsClosed := make(chan bool)

	//we are calling two normal functions using go routine
	go readChat(chChat)
	go sendMessage(chChat, chChatIsClosed)

	// this is necessary to wait the go routines, channels are blocked as default
	<-chChatIsClosed
	fmt.Println("Channel example is Done")
}
