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
	current := numbers.Front()
	for i := 0; i < numbers.Len(); i++ {
		iteration := current
		current = current.Next()
		curValue := iteration.Value.(int)
		curNextValue := iteration.Value.(int)
		if iteration.Next() != nil {
			curNextValue = iteration.Next().Value.(int)
		} else {
			curNextValue = iteration.Value.(int)
		}
		for iteration != nil {
			if (curValue > curNextValue) && (iteration.Next() != nil) {
				iteration = iteration.Next()
				curNextValue = iteration.Value.(int)
			} else if curNextValue > curValue {
				fmt.Print(curNextValue, " ")
				break
			} else if iteration.Next() == nil {
				fmt.Print(0, " ")
				break
			}
		}
	}
}
