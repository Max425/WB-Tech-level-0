package main

import (
	"fmt"
	"strings"
)

func reverseSlice(data []string) string {
	n := len(data)
	for i := 0; i < n/2; i++ {
		invIndex := n - 1 - i
		data[i], data[invIndex] = data[invIndex], data[i]
	}
	return strings.Join(data, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println("Before:", str)
	data := strings.Split(str, " ")
	fmt.Println("After:", reverseSlice(data))
}
