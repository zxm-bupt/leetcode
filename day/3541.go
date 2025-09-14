package day

func maxFreqSum(s string) int {
	hash := make(map[int]int, 26)
	for i := 0; i < 26; i++ {
		hash[i] = 0
	}
	maxVowel := 0
	maxConsonant := 0
	for _, char := range s {
		hash[int(char-'a')]++
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			maxVowel = max(maxVowel, hash[int(char-'a')])
		} else {
			maxConsonant = max(maxConsonant, hash[int(char-'a')])

		}
	}
	return maxVowel + maxConsonant
}
