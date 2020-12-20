package stack

import "errors"

type SequenceStack struct {
	elements []interface{}
}

func NewSequenceStack() *SequenceStack {
	return new(SequenceStack).Init()
}

func (ss *SequenceStack) Init() *SequenceStack {
	ss.elements = make([]interface{}, 0)
	return ss
}

func (ss *SequenceStack) Len() int {
	return len(ss.elements)
}

func (ss *SequenceStack) Empty() bool {
	return ss.Len() == 0
}

func (ss *SequenceStack) Push(data interface{}) {
	ss.elements = append(ss.elements, data)
}

func (ss *SequenceStack) Pop() (data interface{}, err error) {
	if ss.Empty() {
		return nil, errors.New("empty sequence stack")
	}

	data = ss.elements[ss.Len()-1]
	ss.elements = ss.elements[:ss.Len()-1]
	return
}
