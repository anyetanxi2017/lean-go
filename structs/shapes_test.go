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
	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//}
	//
	////长方式面积
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12.0, 6.0}
	//	checkArea(t, rectangle, 72.0)
	//
	//})
	//
	//// 圆形面积
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//
	//})

	// 表格驱动测试方式
	// 列表驱动测试可以成为你工具箱中的得力武器。
	// 如果你要测试一个接口的不同实现，或者传入函数的数据有很多不同的测试需求，这个武器非常给力。
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		//{Rectangle{12,6},72.0},
		//{Circle{10},314.1592653589793},
		// 为了让测试函数更容易理解 可以选择命名下面的域
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
	// 我们可以通过如下命令来运行列表中指定的测试用例： go test -run TestArea/Rectangle
}
