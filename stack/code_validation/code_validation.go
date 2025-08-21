package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
	var lines []string
	var stack Stack[string]
	var indexes Stack[int]
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	j := 1
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			top, _ := stack.Peek()
			switch {
			case string(line[i]) == "(" || string(line[i]) == "{" || string(line[i]) == "[" || string(line[i]) == "<":
				stack.Push(string(line[i]))
				indexes.Push(j)
			case string(line[i]) == ")" && ((top == "(") && stack.IsEmpty() == false):
				stack.Pop()
				indexes.Pop()
			case string(line[i]) == "}" && ((top == "{") && stack.IsEmpty() == false):
				stack.Pop()
				indexes.Pop()
			case string(line[i]) == "]" && ((top == "[") && stack.IsEmpty() == false):
				stack.Pop()
				indexes.Pop()
			case string(line[i]) == ">" && ((top == "<") && stack.IsEmpty() == false):
				stack.Pop()
				indexes.Pop()
			case string(line[i]) == ")" || string(line[i]) == "}" || string(line[i]) == "]" || string(line[i]) == ">":
				fmt.Println(j)
				return
			}
		}
		j++
	}
	if stack.Size() == 0 {
		fmt.Println(-1)
	} else if stack.Size() > 0 {
		fmt.Println(indexes.items[0])
	}
}
