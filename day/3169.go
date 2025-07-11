func countDays(days int, meetings [][]int) int {

	meetings = mergeSlice(meetings)

	res := meetings[0][0] - 1
	for i := 0; i < len(meetings)-1; i++ {
		if meetings[i+1][0] > meetings[i][1] {
			res += meetings[i+1][0] - meetings[i][1] - 1
		}
	}
	if days > meetings[len(meetings)-1][1] {
		res += days - meetings[len(meetings)-1][1]
	}
	return res
}

func mergeSlice(meetings [][]int) [][]int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	res := make([][]int, 0)
	res = append(res, meetings[0])
	root := 0
	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= res[root][1] {
			res[root][1] = max(res[root][1], meetings[i][1])
		} else {
			res = append(res, meetings[i])
			root++
		}
	}
	return res
}