package main

import "fmt"

type stack []int

func (l *stack) push(num int) {
	*l = append(*l, num)
}

func (l *stack) pop() int {
	n := len(*l)
	num := (*l)[n-1]

	*l = (*l)[:n-1]

	return num
}

func main() {
	var myList stack
	fmt.Println(myList)
	list2 := stack{}
	list2 = append(list2, 6, 8, 9)
	list3 := stack{1, 2, 3, 4}
	list2 = append(list2, list3...)
	fmt.Println(list2)

	list4 := stack{}
	list4.push(65)
	list4.push(66)
	list4.push(67)
	list4.push(68)
	list4.pop()
	fmt.Println(list4)

}
