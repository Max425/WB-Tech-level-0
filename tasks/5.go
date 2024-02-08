package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Print("Input N:")
	var N int
	fmt.Scanln(&N)

	dataChannel := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range dataChannel {
			fmt.Println("Received:", num)
		}
	}()
	deadline := time.After(time.Duration(N) * time.Second)
	for i := 1; ; i++ {
		select {
		case <-deadline:
			close(dataChannel)
			wg.Wait()
			fmt.Println("Program finished")
			return
		default:
			dataChannel <- i
		}
	}
}
