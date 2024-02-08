package main

import "fmt"

func binarySearch(data []int, k int) int {
	left, right := 0, len(data)-1
	med := 0

	for left <= right {
		med = (left + right) / 2

		if data[med] == k {
			return med
		}

		if data[med] > k {
			right = med - 1
		} else {
			left = med + 1
		}
	}
	return -1
}

func main() {
	data := []int{-1, 2, 9, 14, 24, 54, 56, 70, 112}
	fmt.Println(binarySearch(data, 56))
}
