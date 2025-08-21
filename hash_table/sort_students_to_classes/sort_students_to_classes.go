package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	classes := make(map[int][]string)
	for i := 0; i < n; i++ {
		var input string
		_ = scanner.Scan()
		input = scanner.Text()
		parts := strings.Split(input, " ")
		class, _ := strconv.Atoi(parts[0])
		classes[class] = append(classes[class], parts[1])
	}
	for _, class := range []int{9, 10, 11} {
		for _, student := range classes[class] {
			fmt.Printf("%d %s\n", class, student)
		}
	}
}
