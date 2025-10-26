package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			count++
		}
	}

	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			slice[j] = slice[i]
			j++
		}
	}

	*ptr = slice[:count]

	return count
}
