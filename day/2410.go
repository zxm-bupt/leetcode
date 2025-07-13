package day

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(trainers)
	sort.Slice(players, func(i, j int) bool {
		return players[i] >= players[j]
	})

	count := 0

	for _, player := range players {
		n := len(trainers)
		if n == 0 {
			break
		}
		index := sort.Search(n, func(i int) bool {
			return trainers[i] >= player
		})
		if index != n {
			trainers = append(trainers[:index], trainers[index+1:]...)
			count++
		}
	}
	return count
}
