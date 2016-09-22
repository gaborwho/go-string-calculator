package add

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	stringNumberArray := split(numbers)
	integerNumbers := parse(stringNumberArray)
	return sum(integerNumbers)
}

func split(numbers string) []string {
	return strings.Split(numbers, ",")
}

func parse(numbers []string) []int {
	integerNumbers := make([]int, len(numbers))

	for i, numberString := range numbers {
		integerNumbers[i], _ = strconv.Atoi(numberString)
	}

	return integerNumbers
}

func sum(numbers []int) int {
	if len(numbers) > 1 {
		return numbers[0] + numbers[1]
	}

	return numbers[0]
}
