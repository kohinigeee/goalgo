package goalgo

import "fmt"

type AlgoStruct struct {
	Value int
}

func NewAlgoStruct() *AlgoStruct {
	return &AlgoStruct{Value: 10}
}

func HelloWorld() {
	fmt.Println("Hello World")
}
