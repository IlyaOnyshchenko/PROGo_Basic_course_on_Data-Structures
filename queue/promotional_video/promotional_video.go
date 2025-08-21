package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 4000000003)
	scanner.Scan()
	str := scanner.Text()
	str1 := strings.Split(str, " ")
	n, _ = strconv.Atoi(string(str1[0]))
	k, _ = strconv.Atoi(string(str1[1]))
	scanner.Scan()
	input := scanner.Text()
	customers := intInput(input)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)
	maxLeft[0] = customers[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = maxInt((maxLeft[i-1]), (customers[i]))
	}
	maxRight[n-1] = customers[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = maxInt((maxRight[i+1]), (customers[i]))
	}
	maxViewers := 0
	for i := 0; i < n-k+1; i++ {
		left := maxLeft[i]
		right := 0
		if i+k < n {
			right = maxRight[i+k]
		}
		maxViewers = maxInt(maxViewers, (left + right))
	}
	fmt.Println(maxViewers)
}
func intInput(str string) []int {
	var result []int
	str1 := strings.Fields(str)
	for i := 0; i < len(str1); i++ {
		val, _ := strconv.Atoi(string(str1[i]))
		result = append(result, val)
	}
	return result
}

func maxInt(number1 int, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}
