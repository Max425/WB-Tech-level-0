package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6"}
	wg.Add(len(keys))
	for i, key := range keys {
		go func(key string, value int) {
			defer wg.Done()
			mutex.Lock()
			data[key] = value
			mutex.Unlock()
			fmt.Printf("Записано: %s -> %d\n", key, value)
		}(key, i+1)
	}
	wg.Wait()

	fmt.Println("Result:", data)
}
