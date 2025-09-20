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
func insertAtFront(head *Node, x int) *Node {
	return &Node{Val: x, Next: head}
}
func main() {
	node := NewNode(42)
	node.Next = NewNode(10)
	node.Next.Next = NewNode(58)
	newNode := insertAtFront(node, 9)
	traverseList(newNode)
}
