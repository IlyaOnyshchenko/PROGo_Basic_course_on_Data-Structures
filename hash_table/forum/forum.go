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
	topicDict := make(map[string][]int, size)
	var input string
	for i := 1; i < size+1; i++ {
		var topic string
		var message int
		scanner.Scan()
		input = scanner.Text()
		if input == "0" {
			scanner.Scan()
			topic = scanner.Text()
			scanner.Scan()
			_ = scanner.Text()
			topicDict[topic] = append(topicDict[topic], i)
		} else {
			message, _ = strconv.Atoi(input)
			scanner.Scan()
			_ = scanner.Text()
			topicDict = messageInsert(topicDict, message, i)
		}
	}
	length := 0
	hotTopic := ""
	for index, value := range topicDict {
		if len(value) > length {
			length = len(value)
			hotTopic = index
		}
	}
	fmt.Println(hotTopic)
}
func messageInsert(hash map[string][]int, message int, index int) map[string][]int {
	for ind, value := range hash {
		for i := 0; i < len(value); i++ {
			if value[i] == message {
				hash[ind] = append(hash[ind], index)
			}
		}
	}
	return hash
}
