package main

import "fmt"

type WindowsFan struct{}

func (wf *WindowsFan) SayWindowsFan() {
	fmt.Println("Yeah! Windows forever!")
}

type LinuxFan struct{}

func (lf *LinuxFan) SayLinuxFan() {
	fmt.Println("Go Linux! Open source power!")
}

// FanAdapter - интерфейс адаптера для фанатов
type FanAdapter interface {
	Say()
}

// WindowsFanAdapter - адаптер для фаната Windows
type WindowsFanAdapter struct {
	windowsFan *WindowsFan
}

func (wfa *WindowsFanAdapter) Say() {
	wfa.windowsFan.SayWindowsFan()
}

func NewWindowsFanAdapter(windowsFan *WindowsFan) FanAdapter {
	return &WindowsFanAdapter{windowsFan}
}

// LinuxFanAdapter - адаптер для фаната Linux
type LinuxFanAdapter struct {
	linuxFan *LinuxFan
}

func (lfa *LinuxFanAdapter) Say() {
	lfa.linuxFan.SayLinuxFan()
}

func NewLinuxFanAdapter(linuxFan *LinuxFan) FanAdapter {
	return &LinuxFanAdapter{linuxFan}
}

func main() {
	myFamily := [2]FanAdapter{NewWindowsFanAdapter(&WindowsFan{}), NewLinuxFanAdapter(&LinuxFan{})}
	for _, member := range myFamily {
		member.Say()
	}
}
