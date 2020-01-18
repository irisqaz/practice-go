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

	r := ""

	for curr := n; curr != nil; curr = curr.next {
		r = r + strconv.Itoa(curr.data) + "->"
	}

	return r
}

func main() {
	n := node{}
	fmt.Println(&n)
	n = node{5, nil}
	fmt.Println(&n)
	n.next = &node{99, nil}
	n.next.next = &node{76, nil}
	fmt.Println(&n)
}
