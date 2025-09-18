package day

import "container/heap"

type TaskManager struct {
	pq *PriorityQueue
}

func Constructor_3408(tasks [][]int) TaskManager {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, task := range tasks {
		t := Task{
			UserId:   task[0],
			TaskId:   task[1],
			Priority: task[2],
		}
		heap.Push(pq, t)
	}
	return TaskManager{
		pq: pq,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}
	heap.Push(this.pq, task)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	for i := 0; i < this.pq.Len(); i++ {
		if (*this.pq)[i].TaskId == taskId {
			(*this.pq)[i].Priority = newPriority
			heap.Fix(this.pq, i)
			break
		}
	}
}

func (this *TaskManager) Rmv(taskId int) {
	for i := 0; i < this.pq.Len(); i++ {
		if (*this.pq)[i].TaskId == taskId {
			heap.Remove(this.pq, i)
			break
		}
	}
}

func (this *TaskManager) ExecTop() int {
	if this.pq.Len() == 0 {
		return -1
	} else {
		x := heap.Pop(this.pq)
		return x.(Task).UserId
	}
}

type Task struct {
	TaskId   int
	UserId   int
	Priority int
}

type PriorityQueue []Task

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Priority > pq[j].Priority {
		return true
	} else if pq[i].Priority == pq[j].Priority {
		return pq[i].TaskId > pq[j].TaskId
	} else {
		return false
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Task))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}
