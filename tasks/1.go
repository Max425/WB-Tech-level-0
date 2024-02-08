package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Println("Hello, my name is", h.Name)
}

type Action struct {
	Human // Встраивание структуры Human в структуру Action
}

func main() {
	action := Action{
		Human: Human{
			Name: "John",
			Age:  30,
		},
	}

	action.Speak() // Этот метод доступен из структуры Action благодаря встраиванию
}
