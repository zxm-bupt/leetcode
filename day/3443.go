package day

func maxDistance(s string, k int) int {
	max, dis := 0, 0
	hashMap := make(map[int]int)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'S':
			hashMap[0]++
		case 'N':
			hashMap[1]++
		case 'E':
			hashMap[2]++
		case 'W':
			hashMap[3]++
		}
		dis = calculateDistance(hashMap, k)
		if dis > max {
			max = dis
		}
	}
	return max

}

func calculateDistance(s map[int]int, k int) int {
	dis := 0
	v0, v1, v2, v3 := s[0], s[1], s[2], s[3]
	if v0 > v1 {
		v0, v1 = v1, v0
	}
	if v2 > v3 {
		v2, v3 = v3, v2
	}
	dis = v1 + v3
	lose := v0 + v2 - 2*k
	if lose < -(v0 + v2) {
		lose = -(v0 + v2)
	}
	dis -= lose
	return dis
}
