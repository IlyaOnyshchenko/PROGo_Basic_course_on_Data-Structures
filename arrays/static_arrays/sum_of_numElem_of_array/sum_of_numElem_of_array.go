package main

import (
	"fmt"
)

func main() {
	var array = [3]int{}
	var sum int
	for i := 0; i < 3; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 3; i++ {
		sum += sumOfNumbers(array[i])
	}
	fmt.Println(sum)
}
func sumOfNumbers(number int) int {
	var summ int
	number = absInt(number)
	for number > 0 {
		summ += number % 10
		number /= 10
	}
	return summ
}
func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
