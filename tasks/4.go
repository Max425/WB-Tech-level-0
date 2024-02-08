package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Print("Input count of workers:")
	var numWorkers int
	fmt.Scanln(&numWorkers)

	dataChannel := make(chan int)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM) //уведомить о нажатии Ctrl+C или о сигнале TERMINATE

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			for data := range dataChannel {
				fmt.Printf("Worker %d: %d\n", id, data)
				time.Sleep(time.Second)
			}
		}(i + 1)
	}

	for i := 1; ; i++ {
		select {
		case <-exit: //при получении сигнала из канала exit закрываем канал ch и ждём завершения всех горутин
			close(dataChannel)
			fmt.Println("Channel is closed")
			wg.Wait()
			return
		default:
			dataChannel <- i
		}
	}
}
