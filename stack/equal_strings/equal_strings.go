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
	var index = 1
	var stack Stack[string]
	var line, word1, word2 string
	var lines []string
	for scanner.Scan() {
		line = scanner.Text()
		if index == 3 {
			break
		}
		lines = append(lines, line)
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "#" {
				stack.Pop()
			} else {
				stack.Push(string(line[i]))
			}
		}
		if index == 1 {
			for i := 0; i < stack.Size(); i++ {
				word1 += stack.items[i]
			}
		} else {
			for i := 0; i < stack.Size(); i++ {
				word2 += stack.items[i]
			}
		}
		index++
		stack.Clear()
	}
	if word1 == word2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
