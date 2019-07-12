package trees

type Tree interface {
	Root() *Node
}

//Node itself is a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n Node) Root() *Node {
	return &n
}

type BST struct {
	root *Node
}

func (t BST) Root() *Node {
	return t.root
}

func Infix(t Tree) []int {
	n := t.Root()
	if n == nil {
		return []int{}
	}

	var ret []int
	if n.Left != nil {
		ret = Infix(n.Left)
	}
	ret = append(ret, n.Value)
	if n.Right != nil {
		ret = append(ret, Infix(n.Right)...)
	}
	return ret
}

func ToBST(arr []int) *BST {
	var root *BST
	if len(arr) == 0 {
		return root
	}

	for _, n := range arr {
		node := &Node{n, nil, nil}
		if root == nil {
			root = &BST{node}
		} else {
			root.insert(node)
		}
	}

	return root
}

func (t BST) insert(n *Node) {
	if t.root.Value >= n.Value {
		if t.root.Left == nil {
			t.root.Left = n
		} else {
			BST{t.root.Left}.insert(n)
		}
	} else {
		if t.root.Right == nil {
			t.root.Right = n
		} else {
			BST{t.root.Right}.insert(n)
		}
	}
}
