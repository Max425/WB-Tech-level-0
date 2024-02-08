package main

import (
	"fmt"
	"strings"
	"time"
)

// При использовании оператора + для конкатенации строк в цикле каждая новая конкатенация создает новую строку,
// что приводит к копированию данных предыдущей строки в новую строку плюс добавлению новых данных. По мере увеличения
// длины строки и числа итераций этот процесс становится все более затратным.
func concatenateWithPlusOperator(stringsToConcat []string) string {
	result := ""
	for _, str := range stringsToConcat {
		result += str
	}
	return result
}

// Функция Join из пакета strings: Эта функция объединяет элементы среза строк в одну строку. Она выполняет одну операцию
// объединения, копируя каждый элемент в результирующую строку. Это гораздо более эффективный подход, чем использование
// оператора +, особенно при конкатенации большого количества строк.
func concatenateWithJoin(stringsToConcat []string) string {
	return strings.Join(stringsToConcat, "")
}

func ranTest(way string, stringsToConcat []string, worker func(s []string) string) {
	start := time.Now()
	result := worker(stringsToConcat)
	elapsed := time.Since(start)

	fmt.Printf("Result length: %d\n", len(result))
	fmt.Printf("Time taken with %s: %s\n", way, elapsed)
}

func getString(length int) []string {
	stringsToConcat := make([]string, length)
	for i := range stringsToConcat {
		stringsToConcat[i] = "test"
	}
	return stringsToConcat
}

func main() {
	length := 10_000
	ranTest("+ operator", getString(length), concatenateWithPlusOperator)
	ranTest("Join", getString(length), concatenateWithJoin)

}
