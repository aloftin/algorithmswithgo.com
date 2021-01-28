package module01

// Sum will sum up all of the numbers passed
// in and return the result

// Original solution
// func Sum(numbers []int) int {
// 	var sum int
// 	for _, val := range numbers {
// 		sum += val
// 	}

// 	return sum
// }

// Recursion solution
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
