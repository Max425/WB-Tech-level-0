package main

import "fmt"

func main() {
	a, b := 10, 15
	fmt.Println(a, b)

	// 1 способ
	//a, b = b, a

	// 2 способ
	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a, b)
}
