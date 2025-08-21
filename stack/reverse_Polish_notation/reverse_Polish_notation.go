package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	_ = scanner.Scan()
	var stack Stack[int]
	line := strings.Split(scanner.Text(), " ")
	for _, value := range line {
		valueStack, _ := strconv.Atoi(value)
		switch {
		case value != "-" && value != "+" && value != "*":
			stack.Push(valueStack)
		case value == "-":
			res := stack.items[stack.Size()-2] - stack.items[stack.Size()-1]
			stack.Pop()
			stack.Pop()
			stack.Push(res)
		case value == "*":
			res := stack.items[stack.Size()-2] * stack.items[stack.Size()-1]
			stack.Pop()
			stack.Pop()
			stack.Push(res)
		case value == "+":
			res := stack.items[stack.Size()-2] + stack.items[stack.Size()-1]
			stack.Pop()
			stack.Pop()
			stack.Push(res)
		}
	}
	fmt.Println(stack.items[0])
}
