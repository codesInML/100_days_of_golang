package perimeter

import (
	"math"
)

type Rectangle struct {
	width float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Perimeter () (value float64) {
	value = 2*(r.width + r.height)
	return
}

func (r Rectangle) Area () (value float64) {
	value = r.width * r.height
	return
}

func (c Circle) Perimeter () (value float64) {
	value = math.Pi * c.radius * 2
	return
}

func (c Circle) Area () (value float64) {
	value = math.Pi * c.radius * c.radius
	return
}

