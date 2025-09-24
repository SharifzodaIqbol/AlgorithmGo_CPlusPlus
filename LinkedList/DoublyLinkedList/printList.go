package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func NewNode(x int) *Node {
	return &Node{
		Val:  x,
		Prev: nil,
		Next: nil,
	}
}
func print(tail *Node) {
	curr := tail
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Prev != nil {
			fmt.Print(" <- ")
		}
		curr = curr.Prev
	}
}
func main() {
	node := NewNode(90)
	node.Next = NewNode(80)
	node.Next.Prev = node
	node.Next.Next = NewNode(70)
	node.Next.Next.Prev = node.Next
	print(node.Next.Next)
}
