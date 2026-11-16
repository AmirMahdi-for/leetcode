package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
func printArea(a Shape) {
	fmt.Println(a.Area())
}
func main() {
	
	rectangle := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 5}

	printArea(rectangle)
	printArea(circle)
}
