package day

func doesAliceWin(s string) bool {
	res := false
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			res = true
			break
		}
	}
	return res
}
