package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func inNode(head *Node, elem int) bool {
	for head != nil {
		if head.Val == elem {
			return true
		}
		head = head.Next
	}
	return false
}

func main() {
	node := &Node{Val: 7, Next: nil}
	node.Next = &Node{Val: 9, Next: nil}
	node.Next.Next = &Node{Val: 18, Next: nil}
	node.Next.Next.Next = &Node{Val: 15, Next: nil}
	fmt.Println(inNode(node, 15))
}
