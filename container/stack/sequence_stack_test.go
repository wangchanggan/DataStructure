package stack

import (
	"fmt"
	"testing"
)

func TestSequenceStack(t *testing.T) {
	ss := NewSequenceStack()
	ss.Push("111")
	ss.Push("222")
	ss.Push("333")
	fmt.Println(ss.Len())
	fmt.Println(ss.Pop())
	fmt.Println(ss)
}