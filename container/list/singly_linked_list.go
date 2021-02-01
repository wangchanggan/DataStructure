package list

import "fmt"

type Element struct {
	next  *Element
	Value interface{}
}

type List struct {
	root *Element
	len  int
}

func (l *List) Init() *List {
	l.root = nil
	l.len = 0
	return l
}

func NewSinglyLinkedList() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) IsEmpty() bool {
	if l.root == nil {
		return true
	}

	return false
}

func (l *List) Append(v interface{}) {
	e := &Element{Value: v}
	if l.IsEmpty() {
		l.root = e
	} else {
		n := l.root
		for n.next != nil {
			n = n.next
		}
		n.next = e
	}
	l.len++
}

func (l *List) RemoveByValue(v interface{}) (count int) {
	if l.IsEmpty() {
		return
	}

	root := l.root
	if root.Value == v {
		l.root = root.next
		l.len--
		count++
	} else {
		for root.next != nil {
			if root.next.Value == v {
				root.next = root.next.next
				l.len--
				count++
			} else {
				root = root.next
			}
		}
	}

	return
}

func (l *List) RemoveByIndex(index int) (v interface{}) {
	if l.IsEmpty() || index < 0 || index >= l.len {
		return
	}

	root := l.root
	if index == 0 {
		l.root = root.next
		l.len--
		return root.Value
	} else {
		for i := 1; i < index; i++ {
			if root.next != nil {
				root = root.next
			} else {
				return
			}
		}

		v = root.next.Value
		root.next = root.next.next
		l.len--
		return
	}

	return nil
}

func (l *List) Show() {
	if !l.IsEmpty() {
		root := l.root
		for {
			fmt.Println(root.Value)
			if root.next != nil {
				root = root.next
			} else {
				break
			}
		}
	}
}
