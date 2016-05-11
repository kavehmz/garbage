package problems

// Node for BST
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert BST
func Insert(root **Node, value int) {
	if *root == nil {
		*root = &Node{Value: value}
		return
	}

	if (*root).Value >= value {
		Insert(&(*root).Right, value)
		return
	}

	Insert(&(*root).Left, value)
	return
}

//BSTCount find
var BSTCount int

// Find BST
func Find(root *Node, value int) (node *Node) {
	BSTCount++
	if root.Value == value {
		return root
	}
	if root.Value <= value {
		return Find(root.Left, value)
	}
	return Find(root.Right, value)
}

// Iterate BST
func Iterate(root *Node, f func(node *Node)) {
	if root.Left != nil {
		Iterate(root.Left, f)
	}
	f(root)
	if root.Right != nil {
		Iterate(root.Right, f)
	}

}
