package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var size, request int
	_ = scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	dictionary := make(map[string]string, size)
	var input string
	for i := 0; i < size; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		line := strings.SplitAfterN(input, " ", 2)
		dictWord := strings.TrimSpace(line[0])
		meaning := line[1]
		dictionary[dictWord] = meaning
	}
	_ = scanner.Scan()
	request, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < request; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		_, ok := dictionary[input]
		if ok == true {
			fmt.Println(dictionary[input])
		} else {
			fmt.Println("Нет в словаре")
		}
	}
}
