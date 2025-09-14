package day

import "slices"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	needUpdateShip := [][]int{}
	for _, ship := range friendships {
		if !hasCommonElement(languages[ship[0]-1], languages[ship[1]-1]) {
			needUpdateShip = append(needUpdateShip, ship)
		}
	}
	count := make([]int, n+1)
	people := make(map[int]struct{}, len(languages)+1)
	for _, ship := range needUpdateShip {
		if _, exist := people[ship[0]]; !exist {
			people[ship[0]] = struct{}{}
			for _, language := range languages[ship[0]-1] {
				count[language]++
			}
		}
		if _, exist := people[ship[1]]; !exist {
			people[ship[1]] = struct{}{}
			for _, language := range languages[ship[1]-1] {
				count[language]++
			}
		}
	}
	num := 0
	for i := 1; i < len(languages)+1; i++ {
		if _, exist := people[i]; exist {
			num++
		}
	}
	minValue := num
	for _, v := range count {
		minValue = min(minValue, num-v)
	}
	return minValue
}

func hasCommonElement(a, b []int) bool {
	for _, element := range a {
		if slices.Contains(b, element) {
			return true
		}
	}
	return false
}

func MinimumTeachings(n int, languages [][]int, friendships [][]int) int {
	return minimumTeachings(n, languages, friendships)
}
