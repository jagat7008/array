package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++

}
func main() {
	mylist := LinkedList{}
	node1 := &node{data: 12}
	node2 := &node{data: 11}
	node3 := &node{data: 10}
	fmt.Println((mylist))
}
