package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6"}
	for i, key := range keys {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			mu.Lock()
			data[key] = value
			mu.Unlock()
			fmt.Printf("Записано: %s -> %d\n", key, value)
		}(key, i+1)
	}
	wg.Wait()

	fmt.Println("Result:", data)
}
