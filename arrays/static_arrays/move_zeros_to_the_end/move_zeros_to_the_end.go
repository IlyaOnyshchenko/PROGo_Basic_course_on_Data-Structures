package main

import (
	"fmt"
)

func main() {
	var array = [5]int{}
	var cnt int
	for i := 0; i < 5; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < 5; i++ {
		if array[i] == 0 {
			cnt++
		}
	}
	j := 0
	for j < 5-cnt {
		for i := 0; i < 5; i++ {
			if array[i] > 0 {
				array[j] = array[i]
				j++
			}
		}
	}
	for j := 4; j > (5-cnt)-1; j-- {
		array[j] = 0
	}
	for _, value := range array {
		fmt.Print(value, " ")
	}
}
