package numbers

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		//numbers := [5]int{1, 2, 3, 4, 5} 声明时不指定长度就是切片 指定则是数组
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d given,%v", got, want, numbers)
		}
	})
}
