package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}
func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return (*s)[len(*s)-1], nil
}
func (s *Stack) Size() int {
	return len(*s)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	line := scanner.Text()
	input := strings.Split(line, " ")
	doks, _ := strconv.Atoi(input[0])
	cargoNums, _ := strconv.Atoi(input[1])
	barrels, _ := strconv.Atoi(input[2])
	stacks := make([]*Stack, cargoNums)
	for i := range stacks {
		stacks[i] = new(Stack)
	}
	maxBarrels := 0
	for i := 0; i < doks; i++ {
		_ = scanner.Scan()
		line = scanner.Text()
		input = strings.Split(line, " ")
		section, _ := strconv.Atoi(input[1])
		gasType, _ := strconv.Atoi(input[2])
		cargoOut, _ := stacks[section-1].Top()
		if input[0] == "+" {
			stacks[section-1].Push(gasType)
		} else if input[0] == "-" && cargoOut == gasType {
			cargoOut, _ = stacks[section-1].Pop()
		} else if input[0] == "-" {
			fmt.Println("Error")
			return
		}
		if sumStacks(stacks) > maxBarrels {
			maxBarrels = sumStacks(stacks)
		}
		if sumStacks(stacks) > barrels {
			fmt.Println("Error")
			return
		}
	}
	if sumStacks(stacks) == 0 {
		fmt.Println(maxBarrels)
	} else {
		fmt.Println("Error")
	}
}
func sumStacks(array []*Stack) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i].Size()
	}
	return sum
}
