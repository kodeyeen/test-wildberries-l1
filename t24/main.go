package main

import (
	"log"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func DistanceBetween(p1 *Point, p2 *Point) float64 {
	// находим гипотенузу по теореме пифагора которая == расстоянию между 2 точками на плоскости
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	d := math.Hypot(dx, dy)

	return d
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) DistanceTo(other *Point) float64 {
	return DistanceBetween(p, other)
}

func main() {
	p1 := NewPoint(6.0, 4.0)
	p2 := NewPoint(13.0, 7.0)

	log.Println(DistanceBetween(p1, p2))
	log.Println(p1.DistanceTo(p2))
}
