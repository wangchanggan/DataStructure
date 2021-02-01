package queue

type SequenceQueue struct {
	ArrayValues []interface{}
}

func NewSequenceQueue() *SequenceQueue {
	return new(SequenceQueue)
}

func (sq *SequenceQueue) IsEmpty() bool {
	if sq == nil || len(sq.ArrayValues) == 0 {
		return true
	}

	return false
}

func (sq *SequenceQueue) Enquene(v interface{}) {
	sq.ArrayValues = append(sq.ArrayValues, v)
}

func (sq *SequenceQueue) Dequene() (v interface{}) {
	if !sq.IsEmpty() {
		v = sq.ArrayValues[0]
		sq.ArrayValues = sq.ArrayValues[1:len(sq.ArrayValues)]
	}

	return
}
