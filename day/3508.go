package day

import "sort"

type DataPackage struct {
	Source      int
	Destination int
	TimeStamp   int
}

type Router struct {
	Queue       []DataPackage
	MemoryLimit int
	Hash        map[DataPackage]struct{}
}

func Constructor_3508(memoryLimit int) Router {
	return Router{
		Queue:       make([]DataPackage, 0),
		MemoryLimit: memoryLimit,
		Hash:        make(map[DataPackage]struct{}),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {

	newDataPackage := DataPackage{
		Source:      source,
		Destination: destination,
		TimeStamp:   timestamp,
	}

	if _, ok := this.Hash[newDataPackage]; ok {
		return false
	}

	if len(this.Queue) == this.MemoryLimit {
		this.Queue = this.Queue[1:]
	}

	this.Queue = append(this.Queue, newDataPackage)
	this.Hash[newDataPackage] = struct{}{}
	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.Queue) == 0 {
		return []int{}
	} else {
		x := this.Queue[0]
		this.Queue = this.Queue[1:]
		delete(this.Hash, x)
		return []int{x.Source, x.Destination, x.TimeStamp}

	}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	count := 0
	start := sort.Search(len(this.Queue), func(i int) bool {
		return this.Queue[i].TimeStamp >= startTime
	})
	end := sort.Search(len(this.Queue), func(i int) bool {
		return this.Queue[i].TimeStamp > endTime
	})

	for i := start; i < end; i++ {
		if this.Queue[i].Destination == destination {
			count++
		}
	}
	return count
}
