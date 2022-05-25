package homework

func reverse(input []int64) (result []int64) {

	reverseSlice := make([]int64, len(input), len(input))

	for i := len(input)-1; i >= 0; i-- {
		reverseSlice[len(input) - i - 1] = input[i]
	}

	return reverseSlice
}
