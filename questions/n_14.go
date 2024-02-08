package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) { // как и в 13 номере - работаем с копией
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice) // [b b a]
	}(slice)
	fmt.Print(slice) // [a a]
}
