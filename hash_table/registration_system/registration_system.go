package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var size int
	_ = scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	hashTable1 := make(map[string]int, size)
	var input string
	for i := 0; i < size; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		_, ok := hashTable1[input]
		if ok == false {
			hashTable1[input] += 1
			fmt.Println("OK")
		} else {
			for j := 1; j > 0; j++ {
				_, ok = hashTable1[fmt.Sprintf("%s%d", input, j)]
				if ok == false {
					hashTable1[fmt.Sprintf("%s%d", input, j)] += 1
					fmt.Println(fmt.Sprintf("%s%d", input, j))
					break
				} else {
					continue
				}
			}
		}
	}
}
func classFill(hash map[int][]string) map[int][]string {
	for i := 9; i < 12; i++ {
		hash[i] = append(hash[i])
	}
	return hash
}
func printHashItems(hash map[int][]string) {
	for index, value := range hash {
		for i := 0; i < len(value); i++ {
			fmt.Println(index, value[i])
		}
	}
}
