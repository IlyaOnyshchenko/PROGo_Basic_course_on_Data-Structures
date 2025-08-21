package main

import (
	"bufio"
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
func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}
func (s *Stack[T]) Pop() T {
	var val T
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}
func (s *Stack[T]) Clear() {
	for s.Size() != 0 {
		s.Pop()
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stack Stack[int]
	var lines []string
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	for i := 0; i < len(lines); i++ {
		switch {
		case strings.Contains(lines[i], "push") == true:
			stack.Push(strDec(lines[i]))
			fmt.Println("ok")
		case lines[i] == "pop":
			fmt.Println(stack.Pop())
		case lines[i] == "peek":
			fmt.Println(stack.Peek())
		case lines[i] == "count":
			fmt.Println(stack.Size())
		case lines[i] == "clear":
			fmt.Println("ok")
			stack.Clear()
		case lines[i] == "exit":
			fmt.Println("bye")
			return
		}
	}
}
func strDec(str string) int {
	var array []int
	var stepen = 1
	var res int
	for i := 0; i < len(str); i++ {
		elem, err := strconv.Atoi(string(str[i]))
		if err == nil {
			array = append(array, elem)
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		res += array[i] * stepen
		stepen *= 10
	}
	return res
}
