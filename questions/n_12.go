package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1 // инициализация в другом блоке (другая область видимости), выйдя из него мы очищаем эту n
		n++
	}
	fmt.Println(n)
}
