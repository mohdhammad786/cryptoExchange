package datatstructures

type Pair struct {
	Key   any
	Value any
}

type RedBlackTreeComparator[K any] interface {
	Compare(a, b *K) int
}

type TreeNode[K any, V any] struct {
	key   *K
	val   *V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
}

func (treeNode *TreeNode[K, V]) GetKey() *K {
	return treeNode.key
}
func (treeNode *TreeNode[K, V]) GetVal() *V {
	return treeNode.val
}
func (treeNode *TreeNode[K, V]) GetLeft() *TreeNode[K, V] {
	return treeNode.left
}
func (treeNode *TreeNode[K, V]) GetRight() *TreeNode[K, V] {
	return treeNode.right
}

type RedBlackTree[K any, V any] struct {
	root       *TreeNode[K, V]
	comparator RedBlackTreeComparator[K]
}

func NewRedBlackTree[K any, V any](comparator RedBlackTreeComparator[K]) *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		root:       nil,
		comparator: comparator,
	}
}
func (redBlackTree *RedBlackTree[K, V]) GetRoot() *TreeNode[K, V] {
	var rootCopy = *redBlackTree.root
	return &rootCopy
}
func (redBlackTree *RedBlackTree[K, V]) getLeftChild() *TreeNode[K, V] {
	var leftCopy = *redBlackTree.root.left
	return &leftCopy
}
func (redBlackTree *RedBlackTree[K, V]) getRightChild() *TreeNode[K, V] {
	var rightCopy = *redBlackTree.root.right
	return &rightCopy
}
func (redBlackTree *RedBlackTree[K, V]) Insert(key *K, val *V) {
	if redBlackTree.root == nil {
		redBlackTree.root = &TreeNode[K, V]{key: key, val: val}
		return
	}
	node := redBlackTree.root
	var prev *TreeNode[K, V]
	for node != nil {
		if redBlackTree.comparator.Compare(key, node.key) == -1 {
			prev = node
			node = node.left
		} else if redBlackTree.comparator.Compare(key, node.key) == 1 {
			prev = node
			node = node.right
		} else {
			return
		}
	}
	if redBlackTree.comparator.Compare(key, prev.key) == -1 {
		prev.left = &TreeNode[K, V]{key: key, val: val}
	} else {
		prev.right = &TreeNode[K, V]{key: key, val: val}
	}
}
func (redBlackTree *RedBlackTree[K, V]) Find(key *K) Pair {
	if redBlackTree.root == nil {
		return Pair{Key: nil, Value: nil}
	}
	node := redBlackTree.root
	for node != nil {
		if redBlackTree.comparator.Compare(key, node.key) == -1 {
			node = node.left
		} else if redBlackTree.comparator.Compare(key, node.key) == 1 {
			node = node.right
		} else {
			return Pair{Key: node.key, Value: node.val}
		}
	}
	return Pair{Key: nil, Value: nil}
}
func (redBlackTree *RedBlackTree[K, V]) IsKeyLessThanOrEqualExists(key *K) bool {
	if redBlackTree.root == nil {
		return false
	}
	node := redBlackTree.root
	for node != nil {
		if redBlackTree.comparator.Compare(key, node.key) == -1 {
			node = node.left
		} else if redBlackTree.comparator.Compare(key, node.key) == 1 {
			return true
		} else {
			return true
		}
	}
	return false
}

func (redBlackTree *RedBlackTree[K, V]) InorderTraversal(node *TreeNode[K, V], result *[]*Pair) {
	if node == nil {
		return
	}
	redBlackTree.InorderTraversal(node.left, result)
	*result = append(*result, &Pair{Key: node.key, Value: node.val})
	redBlackTree.InorderTraversal(node.right, result)
}

func (redBlackTree *RedBlackTree[K, V]) PostOrderTraversal(node *TreeNode[K, V], result *[]*Pair) {
	if node == nil {
		return
	}
	redBlackTree.PostOrderTraversal(node.right, result)
	*result = append(*result, &Pair{Key: node.key, Value: node.val})
	redBlackTree.PostOrderTraversal(node.left, result)
}

func (redBlackTree *RedBlackTree[K, V]) GetSortedElements(isSortedAscending bool) []*Pair {
	var sortedElements []*Pair
	if isSortedAscending {
		redBlackTree.InorderTraversal(redBlackTree.root, &sortedElements)
	} else {
		redBlackTree.PostOrderTraversal(redBlackTree.root, &sortedElements)
	}
	return sortedElements
}

func (redBlackTree *RedBlackTree[K, V]) Erase(key *K) int {
	var node = redBlackTree.root
	var prev *TreeNode[K, V]
	for node != nil {
		if redBlackTree.comparator.Compare(key, node.key) == -1 {
			prev = node
			node = node.left
		} else if redBlackTree.comparator.Compare(key, node.key) == 1 {
			prev = node
			node = node.right
		} else {
			break
		}
	}
	if node == nil {
		return -1
	}

	if node.left == nil && node.right != nil {
		rightPart := node.right
		if prev.left == node {
			prev.left = rightPart
		} else {
			prev.right = rightPart
		}
		return 1
	} else if node.right == nil && node.left != nil {
		leftPart := node.left
		if prev.left == node {
			prev.left = leftPart
		} else {
			prev.right = leftPart
		}
		return 1
	}
	redBlackTree.ReArrangeConnectionsAndGetTop(node, prev)
	return 1
}

func (redBlackTree *RedBlackTree[K, V]) ReArrangeConnectionsAndGetTop(node *TreeNode[K, V], grandfather *TreeNode[K, V]) *TreeNode[K, V] {
	prevNode := node
	rightNode := node.right
	root := node
	node = node.left
	for node != nil {
		prevNode = node
		node = node.right
	}
	prevNode.right = rightNode
	if grandfather == nil {
		return root.left
	}
	if grandfather.left == root {
		grandfather.left = root.left
	} else {
		grandfather.right = root.left
	}
	return root.left
}

// type SampleData struct {
// 	rollNumber int
// }

// type UserDefinedRedBlackTreeComparator struct{}

// func (comparator UserDefinedRedBlackTreeComparator) Compare(lhs, rhs *SampleData) int {
// 	if lhs.rollNumber < rhs.rollNumber {
// 		return -1
// 	} else if lhs.rollNumber > rhs.rollNumber {
// 		return 1
// 	}
// 	return 0
// }
