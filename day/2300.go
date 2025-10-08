package day

import "sort"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	return successfulPairs(spells, potions, success)
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, spell := range spells {
		index := sort.Search(len(potions), func(i int) bool {
			return int64(potions[i])*int64(spell) >= success
		})
		res[i] = len(potions) - index
	}
	return res
}
