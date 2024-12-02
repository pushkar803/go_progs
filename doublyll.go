package main

import "fmt"

type Node struct {
	val      string
	prevNode *Node
	nextNode *Node
}

type Dbll struct {
	head *Node
}

func (dll *Dbll) prepend(data string) {

	tempnode := Node{
		val:      data,
		nextNode: dll.head,
	}

	if dll.head == nil {
		dll.head = &tempnode
		return
	}

	dll.head.prevNode = &tempnode
	dll.head = &tempnode

}

func (dll *Dbll) append(data string) {

	tempnode := Node{
		val: data,
	}

	if dll.head == nil {
		dll.head = &tempnode
		return
	}

	currentnode := dll.head
	for currentnode.nextNode != nil {
		currentnode = currentnode.nextNode
	}

	tempnode.prevNode = currentnode
	currentnode.nextNode = &tempnode
}

func (dll *Dbll) print() {

	currentnode := dll.head
	for currentnode != nil {

		var cp, cn string
		if currentnode.prevNode == nil || currentnode.nextNode == nil {
			cp = "nil"
			cn = "nil"
		}

		if currentnode.prevNode != nil {
			cp = currentnode.prevNode.val
		}

		if currentnode.nextNode != nil {
			cn = currentnode.nextNode.val
		}

		fmt.Println(cp + " <---- " + currentnode.val + " ----> " + cn)
		currentnode = currentnode.nextNode
	}

}

func main() {

	var dll Dbll

	x := []string{"1", "2", "3", "4"}

	for _, ele := range x {
		dll.append(ele)
	}
	//dll.prepend("0")

	dll.print()

}
