package numbers

func sumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		// 可以使用语法 `slice[low:high]`获取部分切片。如果在冒号的一侧没有数字就会一直取到最边缘的元素。
		// 我们使用 numbers[1:]取到从索引1到最后一个元素。你可能需要花费一些时间才能熟悉切片操作
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, sum(tail))
		}
	}
	return
}
