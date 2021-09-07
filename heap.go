package structure

import (
	"fmt"
)

type HeapNode struct {
	Value int
	Key   string
}

type Heap struct {
	list   []*HeapNode
	lenght int
}

//插入
func (h *Heap) Insert(one *HeapNode) {
	h.list = append(h.list, one)
	h.lenght = len(h.list) - 1
	h.AdjustHeap(h.lenght)
}

//堆排序
func (h *Heap) Sort(heaps []*HeapNode) {
	length := len(heaps)
	length = length - 1
	if length == 1 {
		return
	}
	if length == 2 {
		h.AdjustHeap(length - 1)
	}
	for length > 0 {
		h.SliceNodeSwap(1, length)
		length--
		h.Heapfiy(length, 1)
	}
	//反序
	minPos := 1
	maxPos := h.lenght
	for minPos < maxPos {
		h.SliceNodeSwap(minPos, maxPos)
		minPos++
		maxPos--
	}
}

//自上而下调整
func (h *Heap) AdjustHeap(length int) {
	if length < 1 {
		return
	}
	if length == 2 {
		if h.list[length].Value > h.list[length-1].Value {
			h.SliceNodeSwap(length, length-1)
		}
		return
	}
	i := length
	for i/2 > 0 && h.list[i].Value > h.list[i/2].Value {
		h.SliceNodeSwap(i, i/2)
		i = i / 2
	}
}

//node slice 交换
func (h *Heap) SliceNodeSwap(i, j int) {
	x := h.list[i]
	h.list[i] = h.list[j]
	h.list[j] = x
}

//自上而下堆化

func (h *Heap) Heapfiy(length, pos int) {
	for {
		maxPos := pos
		if pos*2 < length && h.list[pos].Value < h.list[pos*2].Value {
			maxPos = pos * 2
		}
		if pos*2+1 < length && h.list[maxPos].Value < h.list[pos*2+1].Value {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		h.SliceNodeSwap(pos, maxPos)
		pos = maxPos
	}
}

func (h *Heap) GetTop() *HeapNode {
	if h.lenght == 0 {
		return nil
	}
	top := h.list[1]
	//堆顶与堆底交换
	h.SliceNodeSwap(1, len(h.list)-1)
	length := len(h.list) - 2
	fmt.Println(length)
	h.Heapfiy(length, 1)
	h.list = append(h.list[:length+1], h.list[length+2:]...)
	h.lenght--
	return top
}

//输出heap
func (h *Heap) Show() {
	for one, value := range h.list {
		fmt.Println(one, value)
	}
}
