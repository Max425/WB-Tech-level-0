package main

import (
	"fmt"
	"math/rand"
)

func generateData(length int) []int {
	var data []int
	for i := 0; i < length; i++ {
		data = append(data, rand.Intn(1000))
	}
	return data
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := partition(arr)

	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	left, right := 0, len(arr)-1
	pivot := arr[right]

	for i := 0; i < right; i++ {
		if arr[i] <= pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func main() {
	data := generateData(30)
	fmt.Println(data)
	quickSort(data)
	fmt.Println(data)
}
