package main

// Stack structure
type Stack struct{
	items []int // I'm using slices, but we can use anything else as well. Or even a custom data structure.
}

// Push an item in stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop an item from stack
func (s *Stack) Pop() int {
	length := len(s.items)

	if length == 0 { // return -1 if stack gets empty
		return -1
	}

	item, items := s.items[length-1], s.items[0:length-1] // Beware, indexing starts from 0
	s.items = items
	return item 
	
}
func main() {
	stack :=  &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(8)

	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
}
