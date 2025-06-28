package day

import (
	"sort"
)

type Pair struct {
	Val int
	Idx int
}

func MaxSubsequence(nums []int, k int) []int {
	q := make([]Pair, 0, len(nums))
	for i, v := range nums {
		q = append(q, Pair{Val: v, Idx: i})
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].Val > q[j].Val
	})
	sort.Slice(q[:k], func(i, j int) bool {
		return q[i].Idx < q[j].Idx
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++{
		res = append(res, q[i].Val)
	}
	return res
}	