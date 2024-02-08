package main

import "fmt"

type WindowsFan struct{}

func (wf *WindowsFan) CheerWindowsFan() {
	fmt.Println("Yeah! Windows forever!")
}

type LinuxFan struct{}

func (lf *LinuxFan) CheerLinuxFan() {
	fmt.Println("Go Linux! Open source power!")
}

// FanAdapter - интерфейс адаптера для фанатов
type FanAdapter interface {
	Cheer()
}

// WindowsFanAdapter - адаптер для фаната Windows
type WindowsFanAdapter struct {
	windowsFan *WindowsFan
}

func (wfa *WindowsFanAdapter) Cheer() {
	wfa.windowsFan.CheerWindowsFan()
}

func NewWindowsFanAdapter(windowsFan *WindowsFan) FanAdapter {
	return &WindowsFanAdapter{windowsFan}
}

// LinuxFanAdapter - адаптер для фаната Linux
type LinuxFanAdapter struct {
	linuxFan *LinuxFan
}

func (lfa *LinuxFanAdapter) Cheer() {
	lfa.linuxFan.CheerLinuxFan()
}

func NewLinuxFanAdapter(linuxFan *LinuxFan) FanAdapter {
	return &LinuxFanAdapter{linuxFan}
}

func main() {
	myFamily := [2]FanAdapter{NewWindowsFanAdapter(&WindowsFan{}), NewLinuxFanAdapter(&LinuxFan{})}
	for _, member := range myFamily {
		member.Cheer()
	}
}
