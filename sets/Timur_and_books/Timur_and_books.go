package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
	// "errors"
)

type SetOnMap[T comparable] map[T]struct{}

func NewSetOnMap[T comparable]() SetOnMap[T] {
	return SetOnMap[T]{}
}
func (s SetOnMap[T]) NewDifference(otherSet SetOnMap[T]) {
	for item, _ := range otherSet {
		if _, ok := s[item]; ok {
			delete(s, item)
			delete(otherSet, item)
		}
	}
}
func (s SetOnMap[T]) Print() string {
	result := ""
	for key := range s {
		result += fmt.Sprintf("%v ", key)
	}
	return result
}
func (s SetOnMap[T]) Empty() bool {
	return len(s) == 0
}
func (s SetOnMap[T]) Add(item T) bool {
	if _, ok := s[item]; ok {
		return false
	}
	s[item] = struct{}{}
	return true
}
func (s SetOnMap[T]) Remove(item T) bool {
	for i, _ := range s {
		if i == item {
			delete(s, i)
			return true
		}
	}
	return false
}
func (s SetOnMap[T]) CompareBooks(otherSet SetOnMap[T]) {
	for item, _ := range otherSet {
		if _, ok := s[item]; ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
func main() {
	library := NewSetOnMap[string]()
	var input int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())
	for input > 0 {
		scanner.Scan()
		book := scanner.Text()
		library.Add(book)
		input--
	}
	literature := NewSetOnMap[string]()
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())
	for input > 0 {
		scanner.Scan()
		book := scanner.Text()
		literature.Add(book)
		input--
	}
	library.CompareBooks(literature)
}
