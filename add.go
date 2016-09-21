package add

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	stringNumberArray := Split(numbers)
	integerNumbers := Parse(stringNumberArray)
	return Sum(integerNumbers)
}

func Split(numbers string) []string {
	return strings.Split(numbers, ",")
}

func Parse(numbers []string) []int {
	integerNumbers := make([]int, len(numbers))

	for i, numberString := range numbers {
		integerNumbers[i], _ = strconv.Atoi(numberString)
	}

	return integerNumbers
}

func Sum(numbers []int) int {
	if len(numbers) > 1 {
		return numbers[0] + numbers[1]
	}

	return numbers[0]
}
