package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(x, y int) bool { return h[x] < h[y] }
func (h IntHeap) Swap(x, y int)      { h[x], h[y] = h[y], h[x] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	h := &IntHeap{}
	heap.Init(h)

	curSum := 0

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			heap.Push(h, curSum)
			if len(*h) > 3 {
				heap.Pop(h)
			}
			curSum = 0
		} else {
			cal, _ := strconv.Atoi(l)
			curSum += cal
		}
	}

	total := 0
	for _, v := range *h {
		total += v
		println(v)
	}
	fmt.Printf("Total %d\n", total)
}
