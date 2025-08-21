package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	// "strings"
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
func (q *QueueOnCircleSlice) intInput(str string) *QueueOnCircleSlice {
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			continue
		} else {
			val, _ := strconv.Atoi(string(str[i]))
			q.Push(val)
		}
	}
	return q
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	player1 := NewQueueOnCircleSlice()
	player2 := NewQueueOnCircleSlice()
	var line string
	_ = scanner.Scan()
	line = scanner.Text()
	player1.intInput(line)
	_ = scanner.Scan()
	line = scanner.Text()
	player2.intInput(line)
	drunkGame(player1, player2)
}
func drunkGame(q1 *QueueOnCircleSlice, q2 *QueueOnCircleSlice) {
	step := 0
	for i := 0; i < 1000001; i++ {
		top1, _ := q1.Peek()
		top2, _ := q2.Peek()
		switch {
		case top1 == 0 && top2 == 9:
			cardPlayer1, _ := q1.Pop()
			cardPlayer2, _ := q2.Pop()
			q1.Push(cardPlayer1)
			q1.Push(cardPlayer2)
			step++
		case top2 == 0 && top1 == 9:
			cardPlayer1, _ := q1.Pop()
			cardPlayer2, _ := q2.Pop()
			q2.Push(cardPlayer1)
			q2.Push(cardPlayer2)
			step++
		case top1 > top2:
			cardPlayer1, _ := q1.Pop()
			cardPlayer2, _ := q2.Pop()
			q1.Push(cardPlayer1)
			q1.Push(cardPlayer2)
			step++
		case top1 < top2:
			cardPlayer1, _ := q1.Pop()
			cardPlayer2, _ := q2.Pop()
			q2.Push(cardPlayer1)
			q2.Push(cardPlayer2)
			step++
		}
		if q1.size == 0 {
			result := fmt.Sprintf("second %d", step)
			fmt.Println(result)
			return
		} else if q2.size == 0 {
			result := fmt.Sprintf("first %d", step)
			fmt.Println(result)
			return
		}
	}
	fmt.Println("botva")
	return
}
