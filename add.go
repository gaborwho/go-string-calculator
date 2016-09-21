package add

import "strconv"

func Add(numbers string) int {
	return CastToInt(numbers)
}

func CastToInt(number string) int {
	i, _ := strconv.Atoi(number)
	return i
}
