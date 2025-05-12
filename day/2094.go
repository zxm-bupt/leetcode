package day

import "sort"

func FindEvenNumbers(digits []int) []int {
	result := make([]int, 0)
	used := make([]bool, len(digits))
	hash := make(map[int]interface{})
	var dfs func(int, []int)
	dfs = func(num int, path []int) {
		if len(path) == 3 {
			if num%2 == 0 && num >= 100 {
				if _, ok := hash[num]; ok {
					return
				}
				result = append(result, num)
				hash[num] = nil
			}
			return
		}
		for i := 0; i < len(digits); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			dfs(num*10+digits[i], append(path, digits[i]))
			used[i] = false
		}
	}
	dfs(0, []int{})
	// sort
	sort.Ints(result)
	return result
}
