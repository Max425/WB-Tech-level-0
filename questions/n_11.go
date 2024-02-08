package main

import (
	"fmt"
	"sync"
)

//каждая горутина получает копию объекта sync.WaitGroup, и каждая горутина работает с своей собственной копией.
//Как следствие, вызовы методов Add и Done на копии объекта sync.WaitGroup не влияют на исходный объект

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) { // тут ошибка, надо по указателю передавать
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
