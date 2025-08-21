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

	list := list.New()
	for _, v := range strList {
		if v == "" {
			continue
		}
		iVal, _ := strconv.Atoi(v)
		list.PushBack(iVal)
	}
	number := 0
	count := list.Len() - 1
	current := list.Front()
	for current != nil {
		if current.Value == 1 {
			number += numPow(count)
			count--
		} else {
			count--
		}
		current = current.Next()
	}
	fmt.Println(number)
}
func numPow(stepen int) int {
	item := 2
	if stepen == 0 {
		return 1
	}
	for i := 1; i < stepen; i++ {
		item += item
	}
	return item
}
