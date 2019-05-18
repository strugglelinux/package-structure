package structure

import (
	"fmt"
	"testing"
)

func TestDList(t *testing.T) {
	dlist := new(DoublyList)
	v := dlist.Append("a")
	if !v {
		t.Errorf(" Append fail v:%v\n", v)
	} else {
		fmt.Printf(" Append success v:%v\n", v)
	}
	s := dlist.Size()
	if s != 1 {
		t.Errorf(" dlist size fail s:%v\n", s)
	} else {
		fmt.Printf(" dlist size success s:%v\n", s)
	}
	dlist.Append("b")
	dlist.Append("c")
	dlist.Append("d")
	dlist.Append("e")
	dlist.Append("f")
	fmt.Printf(" dlist size is %v\n", dlist.Size())
	fmt.Printf("-----------------------------\n")
	node := dlist.GetHead()
	if node.data != "a" {
		t.Errorf(" node head fail value:%v\n", node.data)
	} else {
		fmt.Printf("node head is %v\n", node.data)
	}
	for node != nil {
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("-----------------------------\n")

	node = dlist.GetTail()
	if node.data != "f" {
		t.Errorf(" node tail fail value:%v\n", node.data)
	} else {
		fmt.Printf("node tail is %v\n", node.data)
	}
	for node != nil {
		fmt.Printf("node value %v\n", node.data)
		node = node.prev
	}
	fmt.Printf("------------remvoe----------------\n")
	node = dlist.GetHead()
	//remove c
	removeNode := node.next.next
	dlist.Remove(removeNode)
	for node != nil {
		if node.data == "c" {
			t.Errorf(" node remove c fail value:%v\n", node.data)
		}
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("------------remvoe head----------------\n")
	node = dlist.GetHead()
	dlist.Remove(node)
	node = dlist.GetTail()
	for node != nil {
		if node.data == "a" {
			t.Errorf(" node remove a fail value:%v\n", node.data)
		}
		fmt.Printf("node value %v\n", node.data)
		node = node.prev
	}
	fmt.Printf("------------remvoe tail----------------\n")
	node = dlist.GetTail()
	dlist.Remove(node)
	node = dlist.GetHead()
	for node != nil {
		if node.data == "f" {
			t.Errorf(" node remove  f fail value:%v\n", node.data)
		}
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("------------insert Head----------------\n")
	node = dlist.GetHead()
	dlist.InsertBefore(node, "g")
	node = dlist.GetHead()
	if node.data != "g" {
		t.Errorf("InsertBefore fail value:%v\n", node.data)
	} else {
		fmt.Printf("InsertBefore  success %v\n", node.data)
	}
	for node != nil {
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("------------insert tail----------------\n")
	node = dlist.GetTail()
	dlist.InsertNext(node, "k")
	node = dlist.GetTail()
	if node.data != "k" {
		t.Errorf("InsertNext fail value:%v\n", node.data)
	} else {
		fmt.Printf("InsertNext success %v\n", node.data)
	}
	node = dlist.GetHead()
	for node != nil {
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("------------insert next----------------\n")
	node = dlist.GetHead()
	dlist.InsertNext(node, "m")
	if node.next.data != "m" {
		t.Errorf("InsertNext fail value:%v\n", node.next.data)
	} else {
		fmt.Printf("InsertNext success %v\n", node.next.data)
	}
	for node != nil {
		fmt.Printf("node value %v\n", node.data)
		node = node.next
	}
	fmt.Printf("------------clear----------------\n")
	dlist.Clear()
	s = dlist.Size()
	if s != 0 {
		t.Errorf("Clear fail size:%v\n", s)
	} else {
		fmt.Printf("Clear success size:%v\n", s)
	}
}
