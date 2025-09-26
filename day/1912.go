package day

import (
	"container/heap"
)

type MovieRentingSystem struct {
	reportQ *PQ
	search  map[int]*PQ
	entries [][]int
}

func Constructor_1912(n int, entries [][]int) MovieRentingSystem {
	reportQ := &PQ{}
	heap.Init(reportQ)
	search := make(map[int]*PQ)
	for _, entriy := range entries {
		movie := Movie{
			Shop:    entriy[0],
			MovieId: entriy[1],
			Price:   entriy[2],
		}
		if _, ok := search[movie.MovieId]; !ok {
			pq := &PQ{}
			heap.Init(pq)
			search[movie.MovieId] = pq
		}
		pq := search[movie.MovieId]
		heap.Push(pq, movie)

	}
	return MovieRentingSystem{
		reportQ: reportQ,
		search:  search,
		entries: entries,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	res := []int{}
	if _, ok := this.search[movie]; !ok {
		return res
	}
	copy := *this.search[movie]
	tempHeap := &copy
	heap.Init(tempHeap)
	for i := 0; i < 5 && tempHeap.Len() > 0; i++ {
		x := heap.Pop(tempHeap)
		res = append(res, x.(Movie).Shop)
	}
	return res

}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	for i, mov := range *this.search[movie] {
		if mov.Shop == shop {
			heap.Push(this.reportQ, mov)
			heap.Remove(this.search[movie], i)
			break
		}
	}
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	for i, mov := range *this.reportQ {
		if mov.MovieId == movie && mov.Shop == shop {
			heap.Push(this.search[movie], mov)
			heap.Remove(this.reportQ, i)
			break
		}
	}
}

func (this *MovieRentingSystem) Report() [][]int {
	// 创建队列副本以避免修改原队列
	copy := *this.reportQ
	tempHeap := &copy
	heap.Init(tempHeap)
	res := [][]int{}
	for i := 0; i < 5 && tempHeap.Len() > 0; i++ {
		x := heap.Pop(tempHeap)
		mov := []int{x.(Movie).Shop, x.(Movie).MovieId}
		res = append(res, mov)
	}
	return res
}

type Movie struct {
	Shop    int
	MovieId int
	Price   int
}

type PQ []Movie

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].Price < pq[j].Price {
		return true
	} else if pq[i].Price == pq[j].Price {
		if pq[i].Shop < pq[j].Shop {
			return true
		} else if pq[i].Shop == pq[j].Shop {
			if pq[i].MovieId < pq[j].MovieId {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	*pq = append(*pq, x.(Movie))
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}
