package main

import (
	"fmt"
)

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
func print(node *Node) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
}
func recursionReverse(head *Node) (x *Node) {
	if head == nil || head.Next == nil {
		return head
	}
	x = recursionReverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return x
}
func main() {
	node := NewNode(6)
	node.Next = NewNode(9)
	node.Next.Next = NewNode(23)
	print(node)
	fmt.Println()
	print(recursionReverse(node))
}
