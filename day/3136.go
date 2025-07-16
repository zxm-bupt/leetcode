package day

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	hash := map[rune]interface{}{
		'a': nil,
		'e': nil,
		'i': nil,
		'o': nil,
		'u': nil,
		'A': nil,
		'E': nil,
		'I': nil,
		'O': nil,
		'U': nil,
	}

	countY := 0
	countF := 0
	for _, r := range word {
		if !('0' <= r && r <= '9' || 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z') {
			return false
		}
		if '0' <= r && r <= '9' {
			continue
		}
		if _, ok := hash[r]; ok {
			countY++
		} else {
			countF++
		}
	}
	if countY == 0 {
		return false
	}
	if countF == 0 {
		return false
	}

	return true
}

func IsValid(word string) bool {
	return isValid(word)
}
