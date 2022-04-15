package main

import (
	"fmt"
	"math"
)

type Point2D struct {
	x, y int
}

// конструктор для точек
func NewPoint(x int, y int) *Point2D {
	return &Point2D{x, y}
}
func dist(x int, y int) int {
	d := x - y
	if d < 0 {
		d = y - x
	}
	return d
}
func (p *Point2D) Distance(po *Point2D) float64 {
	d1 := dist(p.x, po.x)
	d2 := dist(p.y, po.y)
	return math.Sqrt(float64(d1*d1 + d2*d2))
}
func main() {
	var x1, y1, x2, y2 int
	fmt.Println("Input coordinate x1")
	fmt.Scan(&x1)
	fmt.Println("Input coordinate y1")
	fmt.Scan(&y1)
	Point1 := NewPoint(x1, y1)
	fmt.Println("Input coordinate x2")
	fmt.Scan(&x2)
	fmt.Println("Input coordinate y2")
	fmt.Scan(&y2)
	Point2 := NewPoint(x2, y2)
	fmt.Printf("Distance is %v\n", Point1.Distance(Point2))

}
