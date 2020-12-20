package stack

import (
	"fmt"
	"testing"
)

func TestChainStack(t *testing.T) {
	cs := NewChainStack()
	cs.Push("111")
	cs.Push("222")
	cs.Push("333")
	fmt.Println(cs.Len())
	fmt.Println(cs.Pop())
	fmt.Println(cs)
	fmt.Println(cs.Pop())
	fmt.Println(cs)
}