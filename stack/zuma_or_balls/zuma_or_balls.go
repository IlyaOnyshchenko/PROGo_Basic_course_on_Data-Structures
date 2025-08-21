package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
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
	var stack Stack[int]
	var totalBalls int
	var length string
	length1, _ := strconv.Atoi(length)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 200003)
	_ = scanner.Scan()
	length = scanner.Text()
	_ = scanner.Scan()
	line := scanner.Text()
	line1 := strings.Split(line, " ")
	row := 3
	for row >= 3 {
		begin, _ := strconv.Atoi(line1[0])
		stack.Push(begin)
		top, _ := stack.Peek()
		indexBegin := 0
		indexLast := 0
		row = 1
		for i := 1; i < len(line1)-totalBalls; i++ {
			value, _ := strconv.Atoi(line1[i])
			if value == top {
				stack.Push(value)
				indexLast = i
				row++
			} else if row < 3 {
				stack.Clear()
				stack.Push(value)
				top, _ = stack.Peek()
				indexBegin = i
				row = 1
			} else if row >= 3 {
				top = 10
			}
		}
		if stack.Size() >= 3 {
			totalBalls += stack.Size()
			length1 += 0
			stack.Clear()
			slices.Delete(line1, indexBegin, indexLast+1)
		}
	}
	fmt.Println(totalBalls)
}
