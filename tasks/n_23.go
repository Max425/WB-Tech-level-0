package main

import "fmt"

func deleteByIndex(data []int, index int) []int {
	if index >= 0 && index < len(data) {
		return append(data[:index], data[index+1:]...)
	}
	return data
}

func main() {
	data := []int{-1, 2, 9, 14, 24, 54, 56, 70, 112}
	fmt.Println(data)
	fmt.Println(deleteByIndex(data, 6))
}
