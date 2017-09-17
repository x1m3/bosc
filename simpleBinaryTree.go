package bosc

import (
	"errors"

)

func init() {
	register("simple_binary_tree", func() BinarySearchTree { return newSimpleBinaryTree() })
}

type simpleBinaryTree struct {
	root *simpleBinaryTreeNode
}

func newSimpleBinaryTree() *simpleBinaryTree {
	return &simpleBinaryTree{}
}

func (t *simpleBinaryTree) Add(val Comparable) error {
	if t.root == nil {
		t.root = newSimpleBinaryTreeNode(val)
		return nil
	}
	return t.root.Add(val)
}

func (t *simpleBinaryTree) Count() uint64 {
	if t.root==nil {
		return 0
	}
	return t.root.Count()
}

func (t *simpleBinaryTree) Find(val Comparable) (Comparable, error) {
	if t.root == nil {
		return nil, errors.New("Not Found. Tree is empty")
	}
	node, _, err := t.root.Find(val, nil)
	return node.Val(), err
}

func (t *simpleBinaryTree) Remove(val Comparable) (found bool) {
	if t.root == nil {
		return false
	}

	// We are removing the root value.
	if t.root.Val() == val {
		// There is only one node with no childs
		if t.root.left == nil && t.root.right == nil {
			t.root = nil
			return true
		}
		if t.root.left != nil && t.root.right == nil { //Childs only at left side
			t.root = t.root.left
			return true
		}
		if t.root.left == nil && t.root.right != nil { // Childs only ar right side
			t.root = t.root.right
			return true
		}
		// Rest of cases can be handled by node.remove()
	}
	return t.root.Remove(val)
}

func (t *simpleBinaryTree) Range(valFrom Comparable, valTo Comparable, fn func(val Comparable)) {
	if t.root != nil {
		t.root.Range(valFrom, valTo, fn)
	}
}

func (t *simpleBinaryTree) RangeAll(fn func(val Comparable)) {
	if t.root != nil {
		t.root.RangeAll(fn)
	}
}

func (t *simpleBinaryTree) RangeFrom(val Comparable, fn func(node Comparable)) {
	if t.root != nil {
		t.root.RangeFrom(val, fn)
	}
}

func (t *simpleBinaryTree) RangeTo(val Comparable, fn func(node Comparable)) {
	if t.root != nil {
		t.root.RangeTo(val, fn)
	}
}

func (t *simpleBinaryTree) Min() Comparable {
	if t.root == nil {
		return nil
	} else {
		return t.root.Min()
	}
}

func (t *simpleBinaryTree) Max() Comparable {
	if t.root == nil {
		return nil
	} else {
		return t.root.Max()
	}
}

type simpleBinaryTreeNode struct {
	val   Comparable
	left  *simpleBinaryTreeNode
	right *simpleBinaryTreeNode
}

func newSimpleBinaryTreeNode(val Comparable) *simpleBinaryTreeNode {
	return &simpleBinaryTreeNode{val: val}
}

func (n *simpleBinaryTreeNode) Val() Comparable {
	return n.val
}

func (n *simpleBinaryTreeNode) Add(val Comparable) error {

	compare := val.Compare(n.val)
	switch {
	case compare == 0:
		return errors.New("Duplicate myNumber")
	case compare < 0:
		if n.left != nil {
			return n.left.Add(val)
		} else {
			n.left = newSimpleBinaryTreeNode(val)
		}
	case compare > 0:
		if n.right != nil {
			return n.right.Add(val)
		} else {
			n.right = newSimpleBinaryTreeNode(val)
		}
	}
	return nil
}

func (n *simpleBinaryTreeNode) Count() (count uint64) {
	count=1
	if n.left!=nil {
		count += n.left.Count()
	}
	if n.right!=nil {
		count += n.right.Count()
	}
	return count
}

func (n *simpleBinaryTreeNode) Find(val Comparable, father *simpleBinaryTreeNode) (this *simpleBinaryTreeNode, parent *simpleBinaryTreeNode, err error) {
	switch {
	case val.Compare(n.val) == 0:
		return n, father, nil
	case val.Compare(n.val) < 0:
		if n.left != nil {
			return n.left.Find(val, n)
		}
	case val.Compare(n.val) > 0:
		if n.right != nil {
			return n.right.Find(val, n)
		}
	}
	return nil, father, errors.New("Not Found")
}

func (n *simpleBinaryTreeNode) Remove(val Comparable) (found bool) {

	node, parent, err := n.Find(val, nil)
	if err != nil {
		return false
	}

	// Node has no childs
	if node.left == nil && node.right == nil {
		if parent.left == node {
			parent.left = nil
			return true
		}
		if parent.right == node {
			parent.right = nil
			return true
		}
	}

	// Node has right childs and no left childs
	if node.left == nil && node.right != nil {
		if parent.left == node {
			parent.left = node.right
		} else {
			parent.right = node.right
		}
		return true
	}

	// Node has left childs and no right childs
	if node.left != nil && node.right == nil {
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
		return true
	}
	// Node has left and right childs
	if node.left != nil && node.right != nil {
		min := node.right.Min()
		node.Remove(min)
		node.val = min
		return true
	}

	return false // Never happens
}

func (n *simpleBinaryTreeNode) Range(valFrom Comparable, valTo Comparable, fn func(val Comparable)) {

	if n.left != nil {
		n.left.Range(valFrom, valTo, fn)
	}

	if n.val.Compare(valFrom) >= 0 && n.val.Compare(valTo) <= 0 {
		fn(n.val)
	}

	if n.right != nil {
		n.right.Range(valFrom, valTo, fn)
	}
}

func (n *simpleBinaryTreeNode) RangeAll(fn func(val Comparable)) {
	if n.left != nil {
		n.left.RangeAll(fn)
	}

	fn(n.val)

	if n.right != nil {
		n.right.RangeAll(fn)
	}
}

func (n *simpleBinaryTreeNode) RangeFrom(from Comparable, fn func(node Comparable)) {
	if n.left != nil {
		n.left.RangeFrom(from, fn)
	}

	if n.val.Compare(from) >= 0 {
		fn(n.val)
	}

	if n.right != nil {
		n.right.RangeFrom(from, fn)
	}
}

func (n *simpleBinaryTreeNode) RangeTo(to Comparable, fn func(node Comparable)) {
	if n.left != nil {
		n.left.RangeTo(to, fn)
	}

	if n.val.Compare(to) <= 0 {
		fn(n.val)
	}

	if n.right != nil {
		n.right.RangeTo(to, fn)
	}
}

func (n *simpleBinaryTreeNode) Min() Comparable {
	p := n
	for {
		if p.left==nil {
			return p.val
		}
		p = p.left
	}
}

func (n *simpleBinaryTreeNode) Max() Comparable {
	p := n
	for {
		if p.right==nil {
			return p.val
		}
		p = p.right
	}
}
