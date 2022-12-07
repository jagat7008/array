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

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) deleteWithValue(value int) {
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := LinkedList{}
	node1 := &node{data: 12}
	node2 := &node{data: 11}
	node3 := &node{data: 10}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	//mylist.deleteWithValue(11)
	mylist.printListData()
}
