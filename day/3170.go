package day

func clearStars(s string) string {
	res := ""
	stacks := make([][]int, 26)
	for i := 0; i < len(stacks); i++ {
		stacks[i] = make([]int, 0)
	}
	arr := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(stacks[j]) > 0 {
					index := stacks[j][len(stacks[j])-1]
					stacks[j] = stacks[j][:len(stacks[j])-1]
					arr[index] = '*'
					break
				}
			}
		} else {
			index := s[i] - 'a'
			stacks[index] = append(stacks[index], i)
		}
	}
	for _, c := range arr {
		if c != '*' {
			res += string(c)
		}
	}
	return res
}
