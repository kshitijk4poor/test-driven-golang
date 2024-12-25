package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var tails []int
	for _, number := range numbers {
		if len(number) == 0 {
			tails = append(tails, 0)
		} else {
			tail := number[1:]
			tails = append(tails, Sum(tail))
		}

	}
	return tails
}
