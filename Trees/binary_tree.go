package main

// Tree is a collection of nodes
type Tree struct {
	node *Node
}

// Node is atomic component of tree
type Node struct {
	 value int
	 left *Node
	 right *Node
}

func (t *Tree) insert(value int) *Tree {
	if t.node == nil{
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
	return t
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

// PrintNode func will print the nodes
func PrintNode(n *Node) {
	if n == nil{
		return
	}
	println(n.value)
	PrintNode(n.left)
	PrintNode(n.right)

}
func main(){
	t := &Tree{}
	t.insert(10).
	insert(20).
	insert(30).
	insert(40).
	insert(50)

	PrintNode(t.node)

}
