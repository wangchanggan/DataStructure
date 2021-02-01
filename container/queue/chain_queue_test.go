package queue

import (
	"fmt"
	"testing"
)

func TestChainQueue(t *testing.T) {
	cq := NewChainQueue()
	cq.Enquene(1)
	cq.Enquene(2)
	cq.Enquene(3)
	cq.Show()
	fmt.Println(cq.Dequene())
	cq.Show()
}
