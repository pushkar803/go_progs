package main

import "fmt"

type llNode struct {
	data string
	next *llNode
}

type LinkedList struct {
	head *llNode
}

func (ll *LinkedList) append(data string) {

	newNode := llNode{
		data: data,
	}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode

}

func (ll *LinkedList) find(data string) bool {
	if ll.head == nil {
		return false
	}

	current := ll.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) print() {
	if ll.head == nil {
		fmt.Println("empty ll")
		return
	}

	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}

func (ll *LinkedList) prepend(data string) {
	newNode := llNode{
		data: data,
	}

	if ll.head == nil {
		ll.head = &newNode
		return
	}

	newNode.next = ll.head
	ll.head = &newNode

}

func main11() {
	ll := &LinkedList{}
	ll.append("1")
	ll.append("2")
	ll.append("3")

	ll.prepend("0")
	ll.prepend("-1")

	ll.print()

	fmt.Println(ll.find("2"))
	fmt.Println(ll.find("5"))
}
