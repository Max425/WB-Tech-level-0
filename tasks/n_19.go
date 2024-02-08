package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println("Before:", str)
	fmt.Println("After:", reverseString(str))
}
