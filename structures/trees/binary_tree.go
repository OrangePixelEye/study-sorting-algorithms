package trees

type BinarySearchTree struct {
	Value int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func (tree *BinarySearchTree) Insert(value  int) {
	if tree == nil {
		return
	}
	if value <= tree.Value {
		if tree.Left == nil {
			tree.Left = &BinarySearchTree{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BinarySearchTree{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
}

func (tree *BinarySearchTree) Contains(value int) bool {
	if tree == nil {
		return false
	}
	if value == tree.Value {
		return true
	}
	if value < tree.Value {
		if tree.Left == nil {
			return false
		}
		return tree.Left.Contains(value)
	} else {
		if tree.Right == nil {
			return false
		}
		return tree.Right.Contains(value)
	}
}

func (tree *BinarySearchTree) InOrderTransversal() {
  if tree != nil {
		tree.Left.InOrderTransversal()
		println(tree.Value)
		tree.Right.InOrderTransversal()
	}
}

func (tree *BinarySearchTree) PreOrderTransverasl() {
  if tree != nil {
		println(tree.Value)
		tree.Left.PreOrderTransverasl()
		tree.Right.PreOrderTransverasl()
	}
}

func (tree *BinarySearchTree) PostOrderTransversal() {
	if tree != nil {
		tree.Left.PostOrderTransversal()
		tree.Right.PostOrderTransversal()
		println(tree.Value)
	}
}

func (tree *BinarySearchTree) BreadthFirstSeatch(value int) bool {
	queue := []*BinarySearchTree{tree}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == nil {
			continue
		}
		if current.Value == value {
			return true
		}
		queue = append(queue, current.Left)
		queue = append(queue, current.Right)
	}
	return false
}

func (tree *BinarySearchTree) DepthFirstSearch(value int) bool {
	if tree == nil {
		return false
	}
	if value == tree.Value {
		return true
	}
	if value < tree.Value {
		return tree.Left.DepthFirstSearch(value)
	} else {
		return tree.Right.DepthFirstSearch(value)
	}
}
