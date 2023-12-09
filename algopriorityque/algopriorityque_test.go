package algopriorityque_test

import (
	"testing"

	"github.com/kohinigeee/goalgo/algopriorityque"
)

func TestIntCase(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "Poped Value", want: 1},
		{name: "Poped Value", want: 2},
	}

	pq := algopriorityque.NewAlgoPriorityQue[int](func(x, y *int) bool {
		return *x < *y
	})

	pq.Push(5)
	pq.Push(4)
	pq.Push(1)
	pq.Push(2)

	for _, tc := range tests {
		got := pq.Pop()
		if got != tc.want {
			t.Fatalf("[%v]  want : %v, got : %v", tc.name, tc.want, got)
		}
	}

	pq.Push(3)

	tests = []struct {
		name string
		want int
	}{
		{name: "Poped Value", want: 3},
		{name: "Poped Value", want: 4},
		{name: "Poped Value", want: 5},
	}

	for _, tc := range tests {
		got := pq.Pop()
		if got != tc.want {
			t.Fatalf("[%v]  want : %v, got : %v", tc.name, tc.want, got)
		}
	}
}
