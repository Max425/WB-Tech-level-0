package main

import "fmt"

var justString string // не стоит создавать в глобальном скопе

func createHugeString(length int) string {
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		buffer[i] = 'a'
	}

	return string(buffer)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // проблема в том, что GC не очистит эту переменную
}

func main() {
	someFunc()
	fmt.Println(justString)
	justString = "" // либо надо очистить после использования
}
