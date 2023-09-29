package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y} // возвращаем указатель на новый объект Point
}

func (a *Point) Distance(b *Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2)) // вычисляем и возвращаем расстояние между двумя точками
}

func (p *Point) PrintPoint() {
	fmt.Printf("{%f, %f}\n", p.x, p.y)
}

func main() {
	p1 := NewPoint(3, 4)
	p1.PrintPoint()
	p2 := NewPoint(7, 7)
	p2.PrintPoint()
	fmt.Println("Расстояние между точками:", p1.Distance(p2))
}
