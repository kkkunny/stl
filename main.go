package main

import (
	"fmt"
	"stl/queue"
	"stl/types"
)

func main() {
	q := queue.NewPriorityQueue[types.I32, types.I32]()
	q.Push(10, 10)
	q.Push(20, 20)
	fmt.Println(q)
}
