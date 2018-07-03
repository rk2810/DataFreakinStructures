package main

type Node struct {
	value int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func(l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	if l.head == nil{
		println("Empty list.")
	}
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil{
			l.head = node
	} else{
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}


func (l *List) InMid(value , index int) {
	node := &Node{value: value}
	i := 0
	n := l.First()
	if index == 0 {
		node.next = n
		n.prev = node
		l.head = node
	}
	for{
		if i == index - 1 {
			if n.Next() != nil {
				node.next = n.next
				n.next.prev = node
				node.prev = n
				n.next = node
				break
			} else {
				n.next = node
				node.prev = n
				l.tail = node
				break
			}
		}
	n = n.Next()
	i++
	}
}


func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	// l.InMid(5,1)
	l.InMid(6,3)

	n := l.First()
	for{
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	
	}

	n = l.Last()
	for{
		println(n.value)
		n = n.Prev()
		if n == nil{
			break
		}
	}

}
