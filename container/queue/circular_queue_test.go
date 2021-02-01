package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	cq := NewCircularQueue(3)
	cq.Enquene(1)
	fmt.Println(cq)
	cq.Enquene(2)
	fmt.Println(cq)
	cq.Enquene(3)
	fmt.Println(cq)
	fmt.Println(cq.Dequene())
	fmt.Println(cq)
}
