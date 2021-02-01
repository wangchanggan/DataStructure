package queue

import (
	"fmt"
	"testing"
)

func TestSequenceQueue(t *testing.T) {
	sq := NewSequenceQueue()
	sq.Enquene(1)
	sq.Enquene(2)
	sq.Enquene(3)
	fmt.Println(sq)
	fmt.Println(sq.Dequene())
	fmt.Println(sq)
}
