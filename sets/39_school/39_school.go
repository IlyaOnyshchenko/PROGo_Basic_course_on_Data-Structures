package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	size, _ := strconv.Atoi(scanner.Text())
	nameDict := make(map[string]int, size)
	for size > 0 {
		scanner.Scan()
		input := scanner.Text()
		nameDict[input] += 1
		size--
	}
	var quantity int
	for _, value := range nameDict {
		if value > 1 {
			quantity += value
		}
	}
	fmt.Println(quantity)
}
