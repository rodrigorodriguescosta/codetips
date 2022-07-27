package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	waitGroupExample()
}

func waitGroupExample() {
	maxSeconds := 5
	minSeconds := 1

	worker1 := func(value int, wg *sync.WaitGroup) {
		// this sleeps simulate some assincronous processing
		sleepSeconds := rand.Intn(maxSeconds - minSeconds)
		time.Sleep(time.Duration(sleepSeconds))

		fmt.Println("Working with: ", value)
		// when the proccess is done, we have to say to WaitGroup that is Done
		wg.Done()
	}

	rangeSize := 10
	wg := &sync.WaitGroup{}
	wg.Add(rangeSize)
	for i := 0; i < rangeSize; i++ {
		go worker1(i, wg)
	}

	wg.Wait()
	fmt.Println("Wait group example has been done")
}
