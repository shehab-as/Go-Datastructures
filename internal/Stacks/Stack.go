package Stacks

type Stack struct {
	A 		[]int
	top 	int 
	length 	int
}

func New() *Stack {
	return &Stack{
		A: nil,
		top: -1,
		length: 0,
	}
}

