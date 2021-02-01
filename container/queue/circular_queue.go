package queue

type CircularQueue struct {
	ArrayValues []interface{}
	cap         int
	head        int
	tail        int
}

func NewCircularQueue(cap int) *CircularQueue {
	return new(CircularQueue).Init(cap)
}

func (cq *CircularQueue) Init(cap int) *CircularQueue {
	cq.ArrayValues = make([]interface{}, cap, cap)
	cq.cap = cap
	return cq
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

func (cq *CircularQueue) Enquene(v interface{}) bool {
	if (cq.tail+1)%cq.cap == cq.head {
		return false
	}

	cq.ArrayValues[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.cap
	return true
}

func (cq *CircularQueue) Dequene() (v interface{}) {
	if cq.IsEmpty() {
		return
	}

	v = cq.ArrayValues[cq.head]
	cq.head = (cq.head + 1) % cq.cap
	return
}
