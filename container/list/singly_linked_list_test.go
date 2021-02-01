package list

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	singlyLinkedList := NewSinglyLinkedList()
	singlyLinkedList.Append(1)
	singlyLinkedList.Append(2)
	singlyLinkedList.Append(3)
	singlyLinkedList.Append(2)
	singlyLinkedList.Show()
	fmt.Println()
	//fmt.Println(singlyLinkedList.RemoveByValue(2))
	fmt.Println(singlyLinkedList.RemoveByIndex(1))
	fmt.Println()
	singlyLinkedList.Show()
}
