package piscine

func Abort(a, b, c, d, e int) int {
	// Put all numbers in a slice
	numbers := []int{a, b, c, d, e}

	// Simple bubble sort for 5 elements
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// Return the middle element (3rd element in 0-based indexing)
	return numbers[2]
}
