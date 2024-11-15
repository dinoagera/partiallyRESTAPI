package main

import "fmt"

const pi = 3.14

type Figure interface {
	Area() float32
	Perimetr() float32
}

func (s *Square) Area() float32 {
	return s.side1 * s.side2
}
func (t *Triangle) Area() float32 {
	return t.side * t.height
}
func (c *Circule) Area() float32 {
	return pi * c.radius * c.radius
}
func (s *Square) Perimetr() float32 {
	return 2 * (s.side1 + s.side2)
}
func (t *Triangle) Perimetr() float32 {
	fmt.Println("Не получилось найти :(")
	return 0
}
func (c *Circule) Perimetr() float32 {
	return 2 * pi * c.radius
}

type Square struct {
	side1 float32
	side2 float32
}
type Triangle struct {
	side   float32
	height float32
}
type Circule struct {
	radius float32
}

func FoundForm(f Figure) {
	fmt.Printf("Ответы для %T:\n", f)
	fmt.Println(f.Area())
	fmt.Println(f.Perimetr())
}
func Change(s *Square, side float32) {
	s.side1 = side
}

func NewFigure(shape string, params ...float32) Figure {
	switch shape {
	case "square":
		if len(params) != 2 {
			panic("Square requires 2 parameters: side1 and side2")
		}
		return &Square{side1: params[0], side2: params[1]}
	case "triangle":
		if len(params) != 2 {
			panic("Triangle requires 2 parameters: side and height")
		}
		return &Triangle{side: params[0], height: params[1]}
	case "circle":
		if len(params) != 1 {
			panic("Circle requires 1 parameter: radius")
		}
		return &Circule{radius: params[0]}
	default:
		panic("Unknown shape type")
	}
}

func main() {
	Square1 := NewFigure("square", 12, 13)
	Triangle1 := NewFigure("triangle", 12, 13)
	Circule1 := NewFigure("circle", 12)
	square1 := Square1.(*Square)
	Change(square1, 32)
	FoundForm(Square1)
	FoundForm(Triangle1)
	FoundForm(Circule1)
}
