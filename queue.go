package main

import "fmt"

type QManager struct {
	q []int
}

func (qm *QManager) enq(data int) {
	qm.q = append(qm.q, data)
	qm.print()
}

func (qm *QManager) deq() {
	qm.q = qm.q[1:len(qm.q)]
	qm.print()
}

func (qm *QManager) print() {
	fmt.Println(qm.q)
}

func main22() {

	var qm QManager

	for i := range [10]int{} {
		qm.enq(i)
	}

	qm.deq()
	qm.deq()
	qm.deq()

	//qm.print()

}
