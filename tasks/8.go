package main

import (
	"fmt"
	"math/rand"
)

// Инвертирование бита
//var num = rand.Int63()
//iPosition := 2
//fmt.Printf("number: %d base2: %b\n", num, num)
//num ^= 1 << iPosition
//fmt.Printf("number: %d base2: %b\n", num, num)

func main() {
	num := rand.Int63()
	fmt.Printf("before: %d base2: %b\n", num, num)

	bitIndex := 2 // от 0 до 63
	var bitValue int

	fmt.Printf("Установить %d-й бит в 0 или 1? ", bitIndex)
	fmt.Scanln(&bitValue)

	if bitValue == 1 {
		num |= 1 << bitIndex // Устанавливаем бит в 1
	} else if bitValue == 0 {
		num &^= 1 << bitIndex // Устанавливаем бит в 0
	}

	fmt.Printf("after: %d base2: %b\n", num, num)
}
