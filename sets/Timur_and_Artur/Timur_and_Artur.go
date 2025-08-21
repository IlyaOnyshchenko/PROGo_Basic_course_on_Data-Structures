package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SetOnMap[T comparable] map[T]struct{}

func NewSetOnMap[T comparable]() SetOnMap[T] {
	return SetOnMap[T]{}
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
func main() {
	cities := NewSetOnMap[string]()
	var input int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())
	for input+1 > 0 {
		var city string
		scanner.Scan()
		city = scanner.Text()
		ok := cities.Add(city)
		if ok == false {
			fmt.Println("EXIST")
			return
		}
		input--
	}
	fmt.Println("OK")
}
