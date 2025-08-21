package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scnr := bufio.NewScanner(os.Stdin)
	_ = scnr.Scan()
	strList := strings.Split(scnr.Text(), " ")

	numbers := list.New()
	for _, v := range strList {
		if v == "" {
			continue
		}
		iVal, _ := strconv.Atoi(v)
		numbers.PushBack(iVal)
	}
	lenList := numbers.Len()
	current := numbers.Front()
	for i := 0; i < lenList; i++ {
		iteration := current
		current = current.Next()
		curValue := iteration.Value.(int)
		curNextValue := iteration.Value.(int)
		if iteration.Next() != nil {
			curNextValue = iteration.Next().Value.(int)
		} else {
			curNextValue = 101
		}
		for iteration != nil {
			if (curValue != curNextValue) && (iteration.Next() != nil) {
				iteration = iteration.Next()
				curNextValue = iteration.Value.(int)
			} else if curNextValue == curValue {
				numbers.Remove(iteration)
				break
			} else {
				iteration = nil
			}
		}
	}
	printList(numbers)
}
func printList(l *list.List) {
	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Print(item.Value, " ")
	}
	fmt.Println()
}
