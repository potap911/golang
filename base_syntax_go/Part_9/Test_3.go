package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, e, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, p := range shapes {
		perimeter += p.perimeter()
	}
	return perimeter
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l + w + l + w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	r := Rectangle{0, 0, 10, 10}
	c := Circle{0, 0, 5}

	fmt.Println(totalArea(&r))
	fmt.Println(totalPerimeter(&r))

	fmt.Println(totalArea(&c))
	fmt.Println(totalPerimeter(&c))
}
