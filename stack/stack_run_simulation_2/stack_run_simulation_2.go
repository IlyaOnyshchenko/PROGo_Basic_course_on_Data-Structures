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
			val, err := stack.Pop()
			if err == nil {
				fmt.Println(val)
			} else {
				fmt.Println(err)
			}
		case lines[i] == "peek":
			val, err := stack.Peek()
			if err == nil {
				fmt.Println(val)
			} else {
				fmt.Println(err)
			}
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
	var flag bool
	for i := 0; i < len(str); i++ {
		elem, err := strconv.Atoi(string(str[i]))
		if err == nil {
			array = append(array, elem)
		} else if string(str[i]) == "-" {
			flag = true
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		res += array[i] * stepen
		stepen *= 10
	}
	if flag == true {
		return -res
	}
	return res
}
