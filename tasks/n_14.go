package main

import (
	"fmt"
	"reflect"
)

func getType(value interface{}) string {
	// Получаем тип значения интерфейса
	typ := reflect.TypeOf(value)

	return typ.String()
}

func main() {
	intValue := 42
	strValue := "Hello"
	boolValue := true
	chValue := make(chan int)

	fmt.Println("Type intValue:", getType(intValue))
	fmt.Println("Type strValue:", getType(strValue))
	fmt.Println("Type boolValue:", getType(boolValue))
	fmt.Println("Type chValue:", getType(chValue))
}
