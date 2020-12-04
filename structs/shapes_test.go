package structs

import (
	"testing"
)

// 长方形周长
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		// f 对应 float64 .2表示输出2位小数
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//长方式面积
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	// 圆形面积
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
