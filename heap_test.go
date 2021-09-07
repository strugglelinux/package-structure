package structure

import (
	"testing"
)

func TestHeap(t *testing.T) {
	arrList := []int{1, 2, 11, 3, 7, 8, 4, 5}
	var myHeap Heap
	myHeap.list = append(myHeap.list, &HeapNode{})
	for _, value := range arrList {
		tmp := HeapNode{}
		tmp.Value = value
		myHeap.Insert(&tmp)
	}

	myHeap.Sort(myHeap.list)
	myHeap.Show()

	for {
		node := myHeap.GetTop()
		if node == nil {
			break
		}
	}
	myHeap.Show()

}
