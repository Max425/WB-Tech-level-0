package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Int63()
	iPosition := 2
	fmt.Printf("number: %d base2: %b\n", num, num)
	num ^= 1 << (iPosition - 1)
	fmt.Printf("number: %d base2: %b\n", num, num)
}
