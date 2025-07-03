package day

func kthCharacter(k int) byte {
	word := "a"
	for len(word) < k {
		temp := ""
		for i := 0; i < len(word); i++ {
			temp += string('a' + (word[i]-'a'+1)%26)
		}
		word += temp
	}
	return word[k-1]
}
