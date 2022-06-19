package main

import (
	"fmt"
)

// 고랭 기본 슬라이스를 이용한 큐임

//The := Short variable declaration can only be used inside functions.

func main() {
	queue := make([]int, 0)
	// Push to the queue
	queue = append(queue, 1)
	// Top (just get next element, don't remove it)
	x := queue[0]
	// Discard top element
	queue = queue[1:]
	// Is empty ?
	if len(queue) == 0 {
		fmt.Println("Queue is empty !")
	}
  fmt.Println("x:",x)
}
