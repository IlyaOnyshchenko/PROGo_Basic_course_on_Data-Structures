package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	// "strconv"
	// "strings"
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
		return zero, errors.New("error")
	}
	return s.items[len(s.items)-1], nil
}
func (s *Stack[T]) Pop() (T, error) {
	var val T
	if len(s.items) == 0 {
		return val, errors.New("error")
	}
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}
func (s *Stack[T]) Clear() {
	for s.Size() != 0 {
		s.Pop()
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stack Stack[string]
	var line string
	var line1 string
	top, _ := stack.Peek()
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}
		line1 = line
	}
	for i := 0; i < len(line1); i++ {
		if string(line1[i]) == top {
			continue
		} else {
			stack.Push(string(line1[i]))
		}
		top, _ = stack.Peek()
	}
	for i := 0; i < stack.Size(); i++ {
		fmt.Print(stack.items[i])
	}
}
