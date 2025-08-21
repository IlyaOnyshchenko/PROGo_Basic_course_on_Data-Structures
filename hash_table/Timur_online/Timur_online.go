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
	hashTable1 := make(map[string][]int, size)
	var input string
	for i := 0; i < size; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		line := strings.Split(input, " ")
		telNumber, _ := strconv.Atoi(line[0])
		hashTable1[line[1]] = append(hashTable1[line[1]], telNumber)
	}
	_ = scanner.Scan()
	request, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < request; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		answer, ok := hashTable1[input]
		if ok == false {
			fmt.Println(fmt.Sprintf("Имени %s нет в телефонной книге", input))
		} else {
			fmt.Println(strings.Trim(fmt.Sprint(answer), "[]"))
		}
	}
}
