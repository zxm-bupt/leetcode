package day

import (
	"sort"
)

func numberOfPairs(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		} else if points[i][0] == points[j][0] {
			if points[i][1] > points[j][1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})
	count := 0
	for i := 0; i < len(points)-1; i++ {
		upper := points[i][1]
		for j := i + 1; j < len(points); j++ {
			if points[j][1] <= upper {
				lower := points[j][1]
				flag := true
				for k := i + 1; k < j; k++ {
					if points[k][1] <= upper && lower <= points[k][1] {
						flag = false
						break
					}
				}
				if flag {
					count++
				}
			}
		}
	}
	return count
}
