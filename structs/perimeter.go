package main

import "math"

func main() {}

type Retangule struct {
	width  float64
	height float64
}

func (r *Retangule) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Retangule) Area() float64 {
	return r.width * r.height
}

type Circule struct {
	size float64
}

func (c *Circule) Area() float64 {
	return math.Pi * math.Pow(c.size, 2)
}

type Geoform interface {
	Area() float64
}
