package structs

import (
	"math"
)

type Shape interface {
	Area() float64
}

// 长方形
type Rectangle struct {
	Width  float64
	Height float64
}

// 圆形
type Circle struct {
	// 半径
	Radius float64
}

// 三角形
type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
