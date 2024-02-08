package main

import "fmt"

func main() {
	data := map[int]int{0: 1, 1: 124, 2: 281}
	for k, v := range data {
		fmt.Println(k, v) // порядок рандомный
	}
}
