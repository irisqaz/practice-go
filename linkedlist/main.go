package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.data)
}

type list struct {
	head *node
	tail *node
	size int
}

func (l list) String() string {

	r := "["

	for curr := l.head; curr != nil; curr = curr.next {
		r = r + curr.String()
		if curr.next != nil {
			r = r + ", "
		}
	}

	return r + "]"
}

func (l *list) length() int {
	return l.size
}

func (l *list) addBack(data int) {
	n := node{data, nil}
	if l.length() == 0 {
		l.head = &n
		l.tail = &n
	} else {
		l.tail.next = &n
		l.tail = &n
	}
	l.size++
}

func (l *list) addFront(data int) {
	n := node{data: data}
	if l.length() == 0 {
		l.head = &n
		l.tail = &n
	} else {
		n.next = l.head
		l.head = &n
	}
	l.size++
}

func main() {
	n := node{}
	fmt.Println(&n)
	n = node{5, nil}
	fmt.Println(&n)
	n.next = &node{99, nil}
	n.next.next = &node{76, nil}
	fmt.Println(&n)

	var n2 *node
	fmt.Println("Empty list n2:", n2)

	n2 = &node{1, n2}
	n2 = &node{2, n2}
	n2 = &node{3, n2}
	fmt.Println(n2)

	var l list
	fmt.Println("Empty list l:", l)
	l.head = &node{10, l.head}
	l.head = &node{11, l.head}
	l.head = &node{12, l.head}
	fmt.Println(l)

	var l2 *list = &list{}
	fmt.Println("Empty list l2:", l2)
	l2.addFront(20)
	l2.addFront(21)
	l2.addFront(22)
	l2.addBack(30)
	l2.addBack(31)
	l2.addBack(32)
	fmt.Println("l2:", l2)
	fmt.Println("length:", l2.length())

	fmt.Println("Traverse Iteratively l2")
	travIter(l2.head)
	fmt.Println()

	fmt.Println("Traverse Recursively l2")
	travRec(l2.head)
	fmt.Println()
}

func travIter(n *node) {
	for curr := n; curr != nil; curr = curr.next {
		fmt.Print(curr.data, " ")
	}
}

func travRec(n *node) {
	if n != nil {
		fmt.Print(n.data, " ")
		travRec(n.next)
	}
}
