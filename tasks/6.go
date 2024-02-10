package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker1(wg *sync.WaitGroup, stop chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("Worker1: Stopping...")
			return
		default:
			fmt.Println("Worker1: Working...")
			time.Sleep(time.Second)
		}
	}
}

func worker2(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker2: Stopping...")
			return
		default:
			fmt.Println("Worker2: Working...")
			time.Sleep(time.Second)
		}
	}
}

func worker3(wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("Worker3: Stopping...")
		wg.Done()
	}()
	for {
		fmt.Println("Worker3: Working...")
		time.Sleep(3 * time.Second)
		runtime.Goexit() // return
	}

}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	stop := make(chan struct{})
	go worker1(wg, stop)

	ctx, cancel := context.WithCancel(context.Background())
	go worker2(wg, ctx)

	go worker3(wg)

	time.Sleep(3 * time.Second)

	//stop <- struct{}{}
	close(stop)
	cancel()
	fmt.Println("Main: Stop signal")

	wg.Wait()
}
