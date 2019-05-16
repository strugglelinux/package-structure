package structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack
	b := s.Push("a")
	if !b {
		t.Errorf("push a %v", b)
	}
	s.Push("b")
	s.Push("c")
	s.Push("d")
	s.Push("e")
	s.Push("f")
	s.Push("g")
	s.Push("h")
	l := s.Len()
	if l != 8 {
		t.Errorf(" len is not 8\n")
	} else {
		fmt.Printf("len is %d \n", l)
	}
	value, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error %v\n", err)
	} else {
		fmt.Printf("Pop value %v\n", value)
	}
	c := s.Cap()
	fmt.Printf("cap is %d \n", c)
	e := s.IsEmpty()
	if e {
		t.Errorf("Stack is Empty cap is  %d\n", c)
	} else {
		fmt.Printf("Stack is not empty! e =%v\n", e)
	}
	top := s.Top()
	if top.(string) != "g" {
		t.Errorf("Top is not g!  top =%s\n", top)
	} else {
		fmt.Printf("Top is %s\n", top)
	}
	s.Clear()
	l = s.Len()
	c = s.Cap()
	fmt.Printf("cap is %d \n", c)
	if l != 0 {
		t.Errorf("Clear fail !  l =%d\n", l)
	} else {
		fmt.Printf("Clear success l=%d\n", l)
	}
}

func TestE(t *testing.T) {
	var s Stack
	data := "(a+b)*[c+d]/{m/6}"
	for c, l := range data {
		fmt.Printf("栈元素个数%d ,c=>%v ,l=%v\n", s.Len(), c, string(l))
		if string(l) == "(" || string(l) == "[" || string(l) == "{" {
			s.Push(string(l))
			fmt.Printf("栈元素个数 %d\n", s.Len())
		} else if string(l) == ")" {
			v := s.Top()
			fmt.Printf("栈元素 v= %v\n", v)
			if v == "(" {
				s.Pop()
			}
		} else if string(l) == "]" {
			v := s.Top()
			fmt.Printf("栈元素 v= %v\n", v)
			if v == "[" {
				s.Pop()
			}
		} else if string(l) == "}" {
			v := s.Top()
			fmt.Printf("栈元素 v= %v\n", v)
			if v == "{" {
				s.Pop()
			}
		}
	}
	if s.IsEmpty() {
		fmt.Printf("字符串括号成对")
	} else {
		t.Errorf("字符串括号不成对")
	}
}
