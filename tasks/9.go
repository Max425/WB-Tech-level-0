package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем горутину для чтения данных из канала input и записи результата в канал output
	go func() {
		defer wg.Done()
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	go func() {
		for result := range output {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
