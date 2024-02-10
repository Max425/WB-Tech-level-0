package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// WaitGroup для ожидания завершения всех горутин
	wg := sync.WaitGroup{}

	// Мьютекс для защиты доступа к переменной sum
	mu := sync.Mutex{}

	var sum int
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			mu.Lock()
			sum += x * x
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}
