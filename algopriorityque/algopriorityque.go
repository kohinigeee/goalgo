package algopriorityque

import (
	"container/heap"
	"log"
)

type priorityQueElement[T any] struct {
	Value    *T
	compFunc *func(*T, *T) bool
}

func newPrioirtyQueElement[T any](value *T, compFunc *func(*T, *T) bool) *priorityQueElement[T] {
	return &priorityQueElement[T]{
		Value:    value,
		compFunc: compFunc,
	}
}

type AlgoHeap[T any] []priorityQueElement[T]

func (heap *AlgoHeap[T]) Len() int {
	return len(*heap)
}

func (heap *AlgoHeap[T]) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *AlgoHeap[T]) Less(i, j int) bool {
	ele1 := (*heap)[i]
	ele2 := (*heap)[j]
	return (*ele1.compFunc)(ele1.Value, ele2.Value)
}

func (h *AlgoHeap[T]) Push(ele interface{}) {
	*h = append(*h, ele.(priorityQueElement[T]))
}

func (h *AlgoHeap[T]) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type AlgoPriorityQue[T any] struct {
	heap     AlgoHeap[T]
	compFunc func(*T, *T) bool
}

func NewAlgoPriorityQue[T any](compFunc func(*T, *T) bool) *AlgoPriorityQue[T] {
	if compFunc == nil {
		log.Fatalf("[NewAlgoPriorityQue: Error] compFunc is nill")
	}
	return &AlgoPriorityQue[T]{
		heap:     make(AlgoHeap[T], 0),
		compFunc: compFunc,
	}
}

func (que *AlgoPriorityQue[T]) Len() int {
	return que.heap.Len()
}

func (que *AlgoPriorityQue[T]) Push(x T) {
	value := x
	ele := newPrioirtyQueElement(&value, &que.compFunc)
	heap.Push(&que.heap, *ele)
}

func (que *AlgoPriorityQue[T]) Pop() T {
	popedValue := heap.Pop(&que.heap).(priorityQueElement[T]).Value
	return *popedValue
}
