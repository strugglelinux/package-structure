package structure

import (
	"errors"
	"sync"
)

//Stack 栈结构
type Stack struct {
	mu    sync.RWMutex
	items []interface{}
}

//init 初始化栈
func (s *Stack) init() {
	s.items = make([]interface{}, 0)
}

//Push 入栈
func (s *Stack) Push(value interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, value)
	return true
}

//Pop 出栈
func (s *Stack) Pop() (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	lenght := len(s.items)
	if lenght == 0 {
		return nil, errors.New("len is 0")
	}
	value := s.items[lenght-1]
	s.items = s.items[:lenght-1]
	return value, nil
}

//Clear 清空栈
func (s *Stack) Clear() {
	s.init()
}

//IsEmpty 判断是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

//Len 返回长度
func (s *Stack) Len() int {
	return len(s.items)
}

//Cap 返回容量
func (s *Stack) Cap() int {
	return cap(s.items)
}

//Top 获取顶元素
func (s *Stack) Top() interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.Len() == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}
