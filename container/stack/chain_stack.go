package stack

type Element struct {
	next  *Element
	Value interface{}
}

type ChainStack struct {
	top *Element
	len int
}

func (cs *ChainStack) Init() *ChainStack {
	cs.top = nil
	cs.len = 0
	return cs
}

func NewChainStack() *ChainStack {
	return new(ChainStack).Init()
}

func (cs *ChainStack) Len() int {
	return cs.len
}

func (cs *ChainStack) Empty() bool {
	return cs.len == 0
}

func (cs *ChainStack) Push(v interface{}) {
	element := new(Element)
	element.Value = v
	element.next = cs.top
	cs.top = element
	cs.len++
}

func (cs *ChainStack) Pop() (v interface{}) {
	if !cs.Empty() {
		v = cs.top.Value
		node := cs.top
		cs.top = node.next
		cs.len--
	}

	return
}
