package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MaxCalorieCounting() int {
	path, err := filepath.Abs("calorie-counting/inventory")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	maxCount := 0
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			if count > maxCount {
				maxCount = count
			}
			count = 0
		} else {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			count += intValue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return maxCount
}

func TopCalorieCounting(k int) int {
	path, err := filepath.Abs("calorie-counting/inventory")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			heap.Push(maxHeap, count)
			count = 0
		} else {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			count += intValue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	maxCount := 0
	for k > 0 {
		value := heap.Pop(maxHeap).(int)
		maxCount += value
		k -= 1
	}
	return maxCount
}

func main() {
	fmt.Println(MaxCalorieCounting())
	fmt.Println(TopCalorieCounting(3))
}
