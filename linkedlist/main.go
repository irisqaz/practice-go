package main

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
}

func (n *node) String() string {

	r := "["

	for curr := n; curr != nil; curr = curr.next {
		r = r + strconv.Itoa(curr.data)
		if curr.next != nil {
			r = r + ", "
		}
	}

	return r + "]"
}

type list struct {
	head *node
	tail *node
}

func (l list) String() string {

	r := "["

	for curr := l.head; curr != nil; curr = curr.next {
		r = r + strconv.Itoa(curr.data)
		if curr.next != nil {
			r = r + ", "
		}
	}

	return r + "]"
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
}
