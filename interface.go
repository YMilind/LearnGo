package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rect struct {
	width  float64
	length float64
}

func (r Rect) Area() float64 {
	return r.width * r.length
}

func main() {

	var s Shape
	s = Rect{5.1, 2.3}
	r := Rect{5.0, 2.0}
	fmt.Printf("%v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Println("%v\n", r.Area())

}
