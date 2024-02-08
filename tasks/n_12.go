package main

import "fmt"

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(strings)

	fmt.Println("Множество:", set)
}
