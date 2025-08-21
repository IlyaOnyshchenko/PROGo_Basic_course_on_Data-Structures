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
	var number int
	scanner.Scan()
	number, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	input := scanner.Text()
	grades := strings.Fields(input)
	gradeDict := make(map[string]int, number*2)
	for i := 0; i < len(grades); i++ {
		gradeDict[grades[i]] += 1
	}
	grade := ""
	for index, value := range gradeDict {
		if value == number {
			grade = index
		}
	}
	total, _ := strconv.Atoi(grade)
	fmt.Println(total)
}
