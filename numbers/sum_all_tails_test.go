package numbers

import (
	"reflect"
	"testing"
)

func TestSumALlTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)

	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)

	})
}

// 总结
// 我们学习了
// 数组
// 切片
//  - 多种方式的切片初始化
//	- 切片的容量是固定的，但可以使用append 从原来的切片中创建一个新切片
//	- 如何获取部分切片
// 使用 len 获取数组和切片的长度
// 使用测试代理覆盖率工具  go test -cover
// reflect.DeepEqual 的用法和对代码类型安全性的影响
