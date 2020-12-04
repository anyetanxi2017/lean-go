package structs

import "math"

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

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
