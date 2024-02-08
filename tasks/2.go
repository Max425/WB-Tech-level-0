package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(num)
	}
	wg.Wait()
}
