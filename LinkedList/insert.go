package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewNode(data int) *Node {
	return &Node{
		Val:  data,
		Next: nil,
	}
}
func insert(head *Node, x int, pos int) *Node {
	if pos < 1 {
		return head
	}
	if pos == 1 {
		return &Node{Val: x, Next: head}
	}
	current := head
	for i := 1; current != nil && i < pos-1; i++ {
		current = current.Next
	}
	if current == nil {
		return head
	}
	newNode := &Node{Val: x, Next: current.Next}
	current.Next = newNode
	return head
}
func traverseList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print("->")
		}
		current = current.Next
	}
}
func main() {
	node := NewNode(18)
	node.Next = NewNode(19)
	node.Next.Next = NewNode(20)
	newNode := insert(node, 2, 0)
	traverseList(newNode)
}
