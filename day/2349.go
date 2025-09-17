package day

import "container/heap"

type NumberContainers struct {
	pq    []*IntHeap
	hash  map[int]int
	array map[int]int
}

func constructor() NumberContainers {
	return NumberContainers{
		pq:    make([]*IntHeap, 0),
		hash:  make(map[int]int),
		array: make(map[int]int),
	}
}

func (this *NumberContainers) change(index int, number int) {
	this.array[index] = number
	if _, ok := this.hash[number]; !ok {
		queue := &IntHeap{}
		heap.Init(queue)
		heap.Push(queue, index)
		flag := len(this.pq)
		this.hash[number] = flag
		this.pq = append(this.pq, queue)
	} else {
		flag := this.hash[number]
		queue := this.pq[flag]
		heap.Push(queue, index)
	}

}

func (this *NumberContainers) find(number int) int {
	if flag, exist := this.hash[number]; exist {
		queue := this.pq[flag]
		for queue.Len() != 0 {
			res := heap.Pop(queue)
			if this.array[res.(int)] == number {
				heap.Push(queue, res)
				return res.(int)
			}
		}
		return -1

	} else {
		return -1
	}
}
