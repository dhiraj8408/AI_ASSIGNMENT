package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type StateNode struct {
	State        string
	ActionPlayed string
	F_N          int64
	G_N          int64
	Parent       *StateNode
	Vertex       int
	Unvisited    *Set[int]
}

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.items[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.items, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.items[value]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Values() []T {
	keys := make([]T, 0, len(s.items))
	for k := range s.items {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set[T]) Copy() *Set[T] {
	newSet := NewSet[T]()
	for k := range s.items {
		newSet.Add(k)
	}
	return newSet
}

type PriorityQueue []*StateNode

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].F_N < pq[j].F_N }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*StateNode)) }
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
func generateGraph(vertices int) [][]int64 {
	adjacencyList := make([][]int64, vertices)
	for i := range adjacencyList {
		adjacencyList[i] = make([]int64, vertices)
		rand.Seed(time.Now().UnixNano())
		for j := range adjacencyList[i] {
			if i == j {
				adjacencyList[i][j] = 0
			} else {
				adjacencyList[i][j] = rand.Int63n(10) + 1
			}
		}
	}
	writeGraphToFile(adjacencyList, "test_case.txt")
	return adjacencyList
}

func writeGraphToFile(graph [][]int64, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "No.of Vertices : %v", len(graph))
	for _, row := range graph {
		for j, val := range row {
			if j != 0 {
				fmt.Fprint(file, " ")
			}
			fmt.Fprint(file, val)
		}
		fmt.Fprintln(file)
	}
}

func GetStateKey(unvisitedVertices *Set[int]) string {
	vertices := unvisitedVertices.Values()
	sort.Ints(vertices)
	return fmt.Sprintf("%v", vertices)
}

type Edge struct {
	cost int64
	to   int
}

type edgeHeap []Edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *edgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}
func (h *edgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
