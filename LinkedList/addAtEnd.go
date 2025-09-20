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

func addAtEnd(head *Node, x int) *Node {
	if head == nil {
		return &Node{Val: x, Next: nil}
	}
	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Node{Val: x, Next: nil}
	return head
}
func printNode(node *Node) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print("->")
		}
		node = node.Next
	}
}
func main() {
	node := NewNode(19)
	node.Next = NewNode(14)
	node.Next.Next = NewNode(155)
	newNode := addAtEnd(node, 999)
	printNode(newNode)
}
