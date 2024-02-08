package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	set := make(map[rune]bool)
	for _, char := range []rune(strings.ToLower(str)) {
		if _, ok := set[char]; ok {
			return false
		}
		set[char] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
}
