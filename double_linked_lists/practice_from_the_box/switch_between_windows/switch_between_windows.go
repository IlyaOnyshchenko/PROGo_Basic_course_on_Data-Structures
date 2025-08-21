package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	appName string
	prev    *Node
	next    *Node
}
type CircularList struct {
	head *Node
	size int
}

func (cl *CircularList) Add(appName string) {
	newNode := &Node{appName: appName}
	if cl.head == nil {
		cl.head = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = cl.head
		newNode.prev = cl.head.prev
		cl.head.prev.next = newNode
		cl.head.prev = newNode
		cl.head = newNode
	}
	cl.size++
}
func (cl *CircularList) Move(k int) {
	if cl.head == nil {
		return
	}
	for i := 0; i < k; i++ {
		cl.head = cl.head.next
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	appList := &CircularList{}
	result := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := scanner.Text()
		if strings.HasPrefix(command, "Run") {
			appName := fmt.Sprintf("%s", command)
			appName = strings.Replace(command, "Run ", "", 1)
			appList.Add(appName)
			result[i] = appName
		} else {
			k := strings.Count(command, "+Tab")
			appList.Move(k)
			result[i] = appList.head.appName
		}
	}
	for _, app := range result {
		fmt.Println(app)
	}
}
