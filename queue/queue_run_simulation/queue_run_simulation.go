package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QueueOnCircleSlice struct {
	items      []int
	size       int
	head, tail int
}

func NewQueueOnCircleSlice() *QueueOnCircleSlice {
	queue := &QueueOnCircleSlice{}
	queue.items = make([]int, 2)
	return queue
}
func (q *QueueOnCircleSlice) Print() string {
	if q.size == 0 {
		return ""
	}
	result := ""
	if q.head < q.tail {
		for i := q.head; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
	} else {
		for i := q.head; i < len(q.items); i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
		for i := 0; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.items[i])
		}
	}
	return result
}
func (q *QueueOnCircleSlice) Empty() bool {
	return q.size == 0
}
func (q *QueueOnCircleSlice) Push(item int) {
	if q.size == cap(q.items) {
		newItems := make([]int, cap(q.items)*2)
		if q.head < q.tail {
			copy(newItems, q.items[q.head:q.tail])
		} else {
			copy(newItems, q.items[q.head:])
			copy(newItems[len(q.items)-q.head:], q.items[:q.tail])
		}

		q.items = newItems
		q.head = 0
		q.tail = q.size
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % cap(q.items)
	q.size++
}
func (q *QueueOnCircleSlice) Pop() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	removed := q.items[q.head]
	q.head = (q.head + 1) % cap(q.items)
	q.size--
	return removed, nil
}
func (q *QueueOnCircleSlice) Peek() (int, error) {
	if q.size == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items[q.head], nil
}
func (q *QueueOnCircleSlice) Clear() {
	q.head = 0
	q.tail = 0
	q.size = 0
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueueOnCircleSlice()
	var lines []string
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	for i := 0; i < len(lines); i++ {
		switch {
		case strings.Contains(lines[i], "enqueue") == true:
			newLine := strings.Split(lines[i], " ")
			number, _ := strconv.Atoi(newLine[1])
			queue.Push(number)
			fmt.Println("ok")
		case lines[i] == "dequeue":
			deleted, _ := queue.Pop()
			fmt.Println(deleted)
		case lines[i] == "peek":
			top, _ := queue.Peek()
			fmt.Println(top)
		case lines[i] == "count":
			fmt.Println(queue.size)
		case lines[i] == "clear":
			fmt.Println("ok")
			queue.Clear()
		case lines[i] == "exit":
			fmt.Println("bye")
			return
		}
	}
}
