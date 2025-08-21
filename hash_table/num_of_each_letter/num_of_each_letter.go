package main

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line)
	hashTable := make(map[string]int, len(line))
	for _, letters := range line {
		_, ok := hashTable[string(letters)]
		if ok == false {
			count := strings.Count(line, string(letters))
			hashTable[string(letters)] = count
		} else {
			continue
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		for index, value := range hashTable {
			if string(i) == index {
				fmt.Println(index, value)
			}
		}
	}
}
