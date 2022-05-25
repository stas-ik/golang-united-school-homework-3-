package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {

	values := make([]string, len(input), len(input))
	keys := make([]int, len(input), len(input))
	counter := 0

	for i := range input {
		keys[counter] = i
		counter++
	}

	sort.Ints(keys)

	for i := 0; i < len(input); i++ {
		values[i] = input[keys[i]]
	}

	return values
}
