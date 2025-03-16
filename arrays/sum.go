package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbersToSum := len(numbersToSum)
	sums := make([]int, lengthOfNumbersToSum) // make slice of capacity lengthOfNumbersToSum

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers) // assign vaue of Sum to slice sums at specified index
	}
	return sums
}
