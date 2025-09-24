package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Val:  data,
		Prev: nil,
		Next: nil,
	}
}
func insertAtegin(node *Node, data int) *Node {
	if node == nil {
		return &Node{Val: data, Prev: nil, Next: nil}
	}
	node.Next.Prev = node.Next
	node = &Node{Val: data, Prev: nil, Next: node}
	return node
}
func print(node *Node) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
}
func main() {
	node := NewNode(18)
	node.Next = NewNode(19)
	node.Next.Prev = node
	node.Next.Next = NewNode(99)
	node.Next.Next.Prev = node.Next
	node = insertAtegin(node, 9)
	print(node)
}
