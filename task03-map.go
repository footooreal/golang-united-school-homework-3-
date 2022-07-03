package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	for i := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, i := range keys {
		result = append(result, input[i])
	}
	return
}
