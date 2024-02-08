package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	distance := Distance(point1, point2)

	fmt.Printf("Расстояние между точкой (%.2f, %.2f) и точкой (%.2f, %.2f) равно %.2f\n",
		point1.x, point1.y, point2.x, point2.y, distance)
}
