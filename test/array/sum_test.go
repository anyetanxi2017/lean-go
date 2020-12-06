package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		// 切片(slice)类型不会将集合的长度保存在类型中，因此 它的尺寸可以是不固定的
		// 切片类型 它可以接收不同大小的集合。语法上和数组非常相似，只是在声明的时间不指定长度
		numbers := []int{1, 2, 3}
		got := sum(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d want %d given,%v", got, want, numbers)
		}

	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
