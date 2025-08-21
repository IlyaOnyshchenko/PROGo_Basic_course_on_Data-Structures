package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(os.Stdin)
	var stack Stack[int]
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if line == "0" {
			break
		}
		lines, _ := strconv.Atoi(line)
		stack.Push(lines)
	}
	for i := stack.Size() - 1; i >= 0; i-- {
		fmt.Print(stack.items[i], " ")
	}
}
