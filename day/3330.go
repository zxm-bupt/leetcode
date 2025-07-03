package day

func possibleStringCount(word string) int {
	if len(word) < 2 {
		return 1
	}
	count := 1
	current := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == current {
			count++
		} else {
			current = word[i]
		}
	}
	return count
}
