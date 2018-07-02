package main

// Node structure stores an integer value and reference to next node
type Node struct {
	value int
	next *Node
}

// List structure has nodes, stores head and tail of list.
type List struct {
	head *Node
	tail *Node
}

// First method returns the head of list.
func(l *List) First() *Node {
	return l.head
}

// Last method returns the last node
func (l *List) Last() *Node{
	if l.head == nil{
		println("Empty list.")
	}
	return l.tail
}

// Search method takes integral value and tells if value exist in list.
func (l *List) Search(search int) {
	if l.head == nil{
		println("Empty list.")
	}
	parameter :=  l.First()
	i := 0
	for {
		if parameter.value == search{
			println("Element exists at", i)
		}
		i++
		parameter = parameter.Next()
		if parameter == l.Last(){
			if parameter.value == search{
				println("Element exists at index", i)
				break
			} else {
			println("Not found")
			break
			}
		}
	}
}

// Push inserts a value into list by creating a node
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil{
			l.head = node
	} else{
		l.tail.next = node
	}

	l.tail = node
}

// Next method returs next node
func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &List{} // instanciate an empty list
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for{
		println(n.value)
		n = n.Next()
		if n == nil{
			break
		}
	
	}

	l.Search(3)

}
