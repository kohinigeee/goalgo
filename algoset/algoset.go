package algoset

type AlgoSet[T comparable] struct {
	values map[T]bool
}

func NewAlgoSet[T comparable](initialValues ...T) *AlgoSet[T] {
	set := &AlgoSet[T]{
		values: make(map[T]bool),
	}

	for _, value := range initialValues {
		set.values[value] = true
	}

	return set
}

func (set *AlgoSet[T]) Insert(insertValues ...T) {
	for _, value := range insertValues {
		set.values[value] = true
	}
}

func (set *AlgoSet[T]) Delete(deleteValues ...T) {
	for _, value := range deleteValues {
		delete(set.values, value)
	}
}

func (set *AlgoSet[T]) Exist(value T) bool {
	_, exist := set.values[value]
	return exist
}
