package main

import "fmt"

type Stack struct {
	items []int
}

// push

func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

//pop
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
