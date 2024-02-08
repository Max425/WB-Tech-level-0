package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewFloat(math.Pow(2, 20) + 10)
	b := big.NewFloat(math.Pow(2, 20) + 15)

	multiplyResult := new(big.Float).Mul(a, b)
	fmt.Println("Умножение:", multiplyResult)

	divideResult := new(big.Float).Quo(a, b)
	fmt.Println("Деление:", divideResult)

	sum := new(big.Float).Add(a, b)
	fmt.Println("Сложение:", sum)

	subtract := new(big.Float).Sub(a, b)
	fmt.Println("Вычитание:", subtract)
}
