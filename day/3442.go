package day

func MaxDifference(s string) int {
	hash := make(map[byte]int, 26)
	for i := 0; i < 26; i++ {
		hash[byte(i)] = 0
	}

	oddMax := 0
	oddMin := 100
	evenMax := 0
	evenMin := 100
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
	}

	for _, v := range hash {
		if v == 0 {
			continue
		}
		if v%2 == 0 {
			evenMax = max(evenMax, v)
			evenMin = min(evenMin, v)
		} else {
			oddMax = max(oddMax, v)
			oddMin = min(oddMin, v)
		}
	}
	return oddMax - evenMin
}
