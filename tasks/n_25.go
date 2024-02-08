package main

import (
	"fmt"
	"time"
)

func sleep(t uint) {
	<-time.After(time.Duration(t) * time.Second)
}

func main() {
	t := 5
	for i := 0; i < t; i++ {
		fmt.Println("Тик")
		sleep(1)
	}
}
