package algoset

import "fmt"

type AlgoSet struct {
	Value int
}

func NewAlgoSet() *AlgoSet {
	return &AlgoSet{Value: 20}
}

func HelloSet() {
	fmt.Println("Hello Set")
}
