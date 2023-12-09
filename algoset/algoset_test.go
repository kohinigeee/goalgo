package algoset_test

import (
	"testing"

	"github.com/kohinigeee/goalgo/algoset"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "Value exists", value: 1, want: true},
		{name: "Value dosen't exist", value: 10, want: false},
	}

	set := algoset.NewAlgoSet[int](1, 2, 3, 4)

	for _, tc := range tests {
		got := set.Exist(tc.value)
		if got != tc.want {
			t.Fatalf("[%v]  want : %v, got : %v", tc.name, tc.want, got)
		}
	}

}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		want     bool
		isInsert bool
	}{
		{name: "Value Insert", value: 1, want: true, isInsert: true},
		{name: "Value don't Insert", value: 10, want: false, isInsert: false},
		{name: "Value dont't Insert and Exist", value: 1, want: true, isInsert: false},
	}

	set := algoset.NewAlgoSet[int]()

	for _, tc := range tests {
		if tc.isInsert {
			set.Insert(tc.value)
		}

		got := set.Exist(tc.value)
		if got != tc.want {
			t.Fatalf("[%v] want : %v, got : %v", tc.name, tc.want, got)
		}
	}
}
