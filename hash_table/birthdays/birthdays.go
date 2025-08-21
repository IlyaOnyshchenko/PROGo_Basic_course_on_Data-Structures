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
	var size int
	_ = scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())
	hashTable1 := make(map[int][]string, size)
	// hashTable1 = classFill(hashTable1)
	var input string
	for i := 0; i < size; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		line := strings.Split(input, " ")
		birthDay, _ := strconv.Atoi(line[2])
		pupil := line[0]
		hashTable1[birthDay] = append(hashTable1[birthDay], pupil)
	}
	_ = scanner.Scan()
	input = scanner.Text()
	request, _ := strconv.Atoi(input)
	for i := 0; i < request; i++ {
		_ = scanner.Scan()
		input = scanner.Text()
		day, _ := strconv.Atoi(input)
		answer, ok := hashTable1[day]
		if ok == false {
			fmt.Println("-")
		} else {
			fmt.Println(strings.Trim(fmt.Sprint(answer), "[]"))
		}
	}
}
