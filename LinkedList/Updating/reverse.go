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
func print(node *Node) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
}
func reverse(head *Node) *Node {
	var prev *Node
	var next *Node
	curr := head
	for curr != nil {
		next = curr.Next // 9 // 23 // nil
		fmt.Println(curr.Next, prev)
		curr.Next = prev // 23 -> 9 -> 6 -> nil
		prev = curr      // 6 // 9 // 23
		curr = next      // 9 // 23 // nil
	}
	return prev
}
func main() {
	node := NewNode(6)
	node.Next = NewNode(9)
	node.Next.Next = NewNode(23)
	print(node)
	fmt.Println()
	print(reverse(node))

}
