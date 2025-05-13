package day

const mod = 1_000_000_007

func LengthAfterTransformations(s string, t int) int {
	hash := make([]int, 26)
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a']++
	}
	for i := 0; i < t; i++ {
		hash = change(hash)
	}
	res := 0
	for _, v := range hash {
		res += v
		res %= mod
	}
	return res
}

func change(hash []int) []int {
	res := make([]int, 26)
	for i := 0; i < 25; i++ {
		res[i+1] = hash[i]
	}
	if hash[25] > 0 {
		res[0] += hash[25]
		res[0] %= mod
		res[1] += hash[25]
		res[1] %= mod
	}
	return res
}
