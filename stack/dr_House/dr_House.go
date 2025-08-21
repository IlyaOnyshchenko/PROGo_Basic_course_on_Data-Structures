package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T int] struct {
	items []int
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack[T]) Push(item int) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() (int, error) {
	var val int
	if len(s.items) == 0 {
		return val, errors.New("error")
	}
	val = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}
func (s *Stack[T]) Peek() (int, error) {
	if s.IsEmpty() {
		var zero int
		return zero, errors.New("error")
	}
	return s.items[len(s.items)-1], nil
}
func (s *Stack[T]) Size() int {
	return len(s.items)
}
func (s *Stack[T]) Clear() {
	s.items = s.items[:1]
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	line := scanner.Text()
	var base Stack[int]
	operation, _ := strconv.Atoi(line)
	for i := 0; i < operation; i++ {
		_ = scanner.Scan()
		line = scanner.Text()
		switch {
		case string(line[0]) == "+":
			input := strings.Split(line, "+")
			gemoLevel, _ := strconv.Atoi(input[1])
			base.Push(gemoLevel)
		case string(line[0]) == "-":
			lastPatient, _ := base.Pop()
			fmt.Println(lastPatient)
		case string(line[0]) == "?":
			input := strings.Split(line, "?")
			patientNum, _ := strconv.Atoi(input[1])
			fmt.Println(base.gemoSum(patientNum))
		}
	}
}
func (s *Stack[T]) gemoSum(num int) int {
	sum := 0
	for i := len(s.items) - 1; i >= len(s.items)-num; i-- {
		sum += s.items[i]
	}
	return sum
}
