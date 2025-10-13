package day

func RemoveAnagrams(words []string) []string {
	return removeAnagrams(words)
}

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	for i := 1; i < len(words); i++ {
		if !isAnagram(words[i-1], words[i]) {
			res = append(res, words[i])
		}
	}

	return res
}

func isAnagram(a, b string) bool {
	hashA := make(map[rune]int)
	hashB := make(map[rune]int)

	for _, char := range []rune(a) {
		hashA[char]++
	}

	for _, char := range []rune(b) {
		hashB[char]++
	}

	if len(hashA) != len(hashB) {
		return false
	}
	for key, v := range hashA {
		if _, ok := hashB[key]; hashB[key] != v || !ok {
			return false
		}
	}
	return true
}
