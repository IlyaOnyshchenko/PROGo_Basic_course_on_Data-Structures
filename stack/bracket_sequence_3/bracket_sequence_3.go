package main

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack[T]) Size() int {
	return len(s.items)
}
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
func (s *Stack[T]) Pop() (T, error) {
	var val T
	if len(s.items) == 0 {
		return val, errors.New("stack is empty")
	}
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func main() {
	var str string
	fmt.Scan(&str)
	var stack Stack[string]
	for i := 0; i < len(str); i++ {
		top, _ := stack.Peek()
		switch {
		case string(str[i]) == "(" || string(str[i]) == "{" || string(str[i]) == "[":
			stack.Push(string(str[i]))
		case string(str[i]) == ")" && ((top == "(") && stack.IsEmpty() == false):
			stack.Pop()
		case string(str[i]) == "}" && ((top == "{") && stack.IsEmpty() == false):
			stack.Pop()
		case string(str[i]) == "]" && ((top == "[") && stack.IsEmpty() == false):
			stack.Pop()
		case string(str[i]) == ")" || string(str[i]) == "}" || string(str[i]) == "]":
			stack.Push(string(str[i]))
		}
	}
	if stack.Size() == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
