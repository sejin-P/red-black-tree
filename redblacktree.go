
type Node interface {
	CheckColor() bool
	CheckIdx() int
	CheckUncleColor() bool
	CheckLeftColor() bool
	CheckRightColor() bool
}

type RealNode struct {
	Left	Node
	Right	Node
	Parent	Node
	IsBlack	bool
	IsLeaf bool
	Idx	int
}

func NewRealNode(idx Idx, isBlack bool) Node {
	return &RealNode{
		IsBlack: isBlack,
		IsLeaf: True,
		Idx: idx
	}
} 

func (realNode RealNode) CheckColor() bool {
	return realNode.IsBlack
}

func (realNode RealNode) CheckIdx() int {
	return realNode.Idx
}

func (realNode RealNode) CheckUncleColor() bool {
	parentNode := realNode.Parent
	if realNode.Idx > parentNode.CheckIdx() {
		return parentNode.CheckLeftColor()
	}
	else {
		return parentNode.CheckRightColor()
	}
}

func (realNode RealNode) CheckLeftColor() bool {
	if realNode.IsLeaf{
		return True
	}
	return realNode.Left.CheckColor()
}

type Tree interface {
	Insert(node Node)
}

type RedBlackTree struct {
	RootNode *RealNode

}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		RootNode: *RealNode{
			Left: nil,
			Right: nil,
			Parent: nil,
			IsBlack: True,
			ISLeaft: True,
			Idx: 0
		}
	}
}

func BinarySearch(rootNode Node, node Node) (parentNode Node, isRight bool) {
	if rootNode.Idx < node.CheckIdx() {
		if rootNode.Right == nil {
			return rootNode, True
		}
		else {
			return rbTree.BinarySearch(rootNode.Right, node)
		}
	}
	else {
		if rootNode.Left == nil {
			return rootNode, False
		}
		else {
			return rbTree.BinarySearch(rootNode.Left, node)
		}
	}
}

func (rbTree RedBlackTree) Insert(node Node) {
	parentNode, isRight := BinarySearch(rbTree.RootNode, node)
	
}

func (rbTree RedBlackTree) Refactoring() {

}

func (rbTree RedBlackTree) Recoloring() {

}
