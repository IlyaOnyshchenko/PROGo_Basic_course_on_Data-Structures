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
	_ = scnr.Scan()
	strList1 := strings.Split(scnr.Text(), " ")

	values := list.New()
	for _, v := range strList1 {
		if v == "" {
			continue
		}
		iVal1, _ := strconv.Atoi(v)
		values.PushBack(iVal1)
	}
	current := numbers.Front()
	current2 := values.Front()
	len1 := 0
	len2 := 0
	for len1 < numbers.Len() || len2 < values.Len() {
		itemNumbers := current.Value.(int)
		itemValues := current2.Value.(int)
		switch {
		case itemNumbers < itemValues && len1 < numbers.Len():
			fmt.Print(itemNumbers, " ")
			if current.Next() != nil {
				current = current.Next()
			}
			len1++
		case itemNumbers > itemValues && len2 < values.Len():
			fmt.Print(itemValues, " ")
			if current2.Next() != nil {
				current2 = current2.Next()
			}
			len2++
		case itemNumbers == itemValues && (len1 < numbers.Len() && len2 < values.Len()):
			fmt.Print(itemValues, " ", itemNumbers, " ")
			if current.Next() != nil {
				current = current.Next()
			}
			if current2.Next() != nil {
				current2 = current2.Next()
			}
			len1++
			len2++
		case len1 == numbers.Len() && len2 <= values.Len():
			fmt.Print(itemValues, " ")
			current2 = current2.Next()
			len2++
		case len2 == values.Len() && len1 <= numbers.Len():
			fmt.Print(itemNumbers, " ")
			current = current.Next()
			len1++
		}
	}
}
