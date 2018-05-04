package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	right *Node
	left  *Node
}

type Tree struct {
	Root *Node
}

func(n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("nil tree, can't add any element")
	} else if value == n.value {
		return errors.New("value already present")
	} else if value < n.value {
		if n.left == nil {
			n.left = &Node{value:value}
			return nil
		} else {
			n.left.Insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = &Node{value:value}
			return nil
		} else {
			return n.right.Insert(value)
		}
	}
	return nil
}

func(n *Node) Find(value int) (int, bool) {
	if n == nil {
		return 0, false
	} else if value == n.value {
		return n.value, true
	} else if value < n.value {
		return n.left.Find(value)
	} else {
		return n.right.Find(value)
	}
}

func(t *Tree) Insert(value int) error {
	if t.Root == nil {
		t.Root = &Node{value:value}
	} else {
		return t.Root.Insert(value)
	}
	return nil
}

func(t *Tree) Find(value int) (int, bool) {
	if t.Root == nil {
		return 0, false
	} else {
		return t.Root.Find(value)
	}
}

func New() *Tree {
	return &Tree{}
}

func main() {

	tree := New()
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)

	fmt.Println(tree.Root.value)
	fmt.Println(tree.Root.left.value)
	fmt.Println(tree.Root.right.value)
	fmt.Println(tree.Root.left.left.value)
	fmt.Println(tree.Root.right.left.value)
	fmt.Println(tree.Root.right.right.value)
	fmt.Println(tree.Find(4))
}