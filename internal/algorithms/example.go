package algorithms

import "math/big"

type Example struct {
	name string
}

func NewExample() *Example {
	return &Example{
		name: "Example",
	}
}

func (e *Example) Name() string {
	return e.name
}

func (e *Example) NextDifficulty() *big.Int {
	return big.NewInt(1)
}
