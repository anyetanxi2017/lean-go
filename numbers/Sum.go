package numbers

func Sum(numbers []int) (sum int) {
	//range 会迭代数组，每次迭代都会返回数组元素的索引和值。
	// 我们选择使用_空白标志符来忽略索引。
	for _, numbers := range numbers {
		sum += numbers
	}
	return
}
