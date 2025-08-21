package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type GoblinQueue struct {
	buffer   []int
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueue(initialCapacity int) *GoblinQueue {
	if initialCapacity < 1 {
		initialCapacity = 1
	}
	return &GoblinQueue{
		buffer:   make([]int, initialCapacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: initialCapacity,
	}
}
func (q *GoblinQueue) Enqueue(goblin int) {
	// if q.size == q.capacity {
	// 	q.resize()
	// }
	if q.buffer[q.tail] == 0 {
		q.buffer[q.tail] = goblin
	} else {
		q.buffer[q.tail+1] = goblin
	}
	q.tail = (q.tail + 1) % q.capacity
	q.size++
}
func (q *GoblinQueue) InsertMiddle(goblin int) {
	// if q.size == q.capacity {
	// 	q.resize()
	// }
	var mid, shiftStart int
	switch {
	case q.size%2 == 0:
		mid = (q.head + (q.size / 2)) % q.capacity
		shiftStart = q.size / 2
	case q.size%2 != 0:
		mid = (q.head + (q.size / 2) + 1) % q.capacity
		shiftStart = q.size/2 + 1
	}
	newItems := make([]int, q.size/2+1)
	copy(newItems, q.buffer[(q.head+shiftStart)%q.capacity:])
	q.buffer[mid] = goblin
	copy(q.buffer[((q.head+shiftStart)%q.capacity)+1:], newItems)
	q.size++
	q.tail = (q.head + q.size) % q.capacity
}

func (q *GoblinQueue) Dequeue() int {
	goblin := q.buffer[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return goblin
}
func (q *GoblinQueue) resize() {
	newCapacity := q.capacity * 2
	newBuffer := make([]int, newCapacity)

	for i := 0; i < q.size; i++ {
		newBuffer[i] = q.buffer[(q.head+i)%q.capacity]
	}

	q.buffer = newBuffer
	q.head = 0
	q.tail = q.size
	q.capacity = newCapacity
}
func (q *GoblinQueue) Print() string {
	if q.size == 0 {
		return ""
	}
	result := ""
	if q.head < q.tail {
		for i := q.head; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.buffer[i])
		}
	} else {
		for i := q.head; i < len(q.buffer); i++ {
			result += fmt.Sprintf("%d ", q.buffer[i])
		}
		for i := 0; i < q.tail; i++ {
			result += fmt.Sprintf("%d ", q.buffer[i])
		}
	}
	return result
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	queue := NewQueue(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		request := scanner.Text()
		switch {
		case request[0] == '+':
			number, _ := strconv.Atoi(request[2:])
			queue.Enqueue(number)
		case request[0] == '*':
			number, _ := strconv.Atoi(request[2:])
			queue.InsertMiddle(number)
		case request[0] == '-':
			fmt.Println(queue.Dequeue())
		}
	}
}
