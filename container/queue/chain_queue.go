package queue

import "fmt"

type Element struct {
	next  *Element
	Value interface{}
}

type ChainQueue struct {
	ele *Element
	len int
}

func (cq *ChainQueue) Init() *ChainQueue {
	cq.ele = nil
	cq.len = 0
	return cq
}

func NewChainQueue() *ChainQueue {
	return new(ChainQueue).Init()
}

func (cq *ChainQueue) Len() int {
	return cq.len
}

func (cq *ChainQueue) IsEmpty() bool {
	return cq.len == 0
}

func (cq *ChainQueue) Enquene(v interface{}) {
	e := new(Element)
	e.Value = v
	if cq.IsEmpty() {
		cq.ele = e
	} else {
		q := cq.ele
		for q.next != nil {
			q = q.next
		}
		q.next = e
	}
	cq.len++
}

func (cq *ChainQueue) Dequene() (v interface{}) {
	if !cq.IsEmpty() {
		v = cq.ele.Value
		cq.ele = cq.ele.next
		cq.len--
	}

	return
}

func (cq *ChainQueue) Show() {
	if !cq.IsEmpty() {
		q := cq.ele
		for q != nil {
			fmt.Println(q.Value)
			q = q.next
		}
	}
}
