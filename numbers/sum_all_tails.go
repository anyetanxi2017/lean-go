package numbers

func SumAll(numberToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numberToSum)
	// make 可以在创建切片的时候 指定需要的长度和容量
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numberToSum {
		sums[i] = sum(numbers)
	}
	return
}
