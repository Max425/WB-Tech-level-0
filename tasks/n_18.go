package main

import (
	"fmt"
	"sync"
)

type ConcurrencyCounter struct {
	mu      sync.Mutex
	counter int
}

func NewConcurrencyCounter() *ConcurrencyCounter {
	return &ConcurrencyCounter{}
}

func (c *ConcurrencyCounter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := NewConcurrencyCounter()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.counter)
}
