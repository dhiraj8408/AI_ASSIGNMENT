package main

import (
	"fmt"
	"time"
)

const MaxInt64 = int64(^uint64(0) >> 1)

// tourCost computes total tour cost for tour represented as slice of vertex indices in order.
// tour should list vertices in visiting order (starting at 0), and tour[0] should be 0.
func tourCost(graph [][]int64, tour []int) int64 {
	n := len(tour)
	if n == 0 {
		return 0
	}
	var cost int64 = 0
	for i := 0; i < n-1; i++ {
		cost += graph[tour[i]][tour[i+1]]
	}
	// return to start
	cost += graph[tour[n-1]][tour[0]]
	return cost
}

// nearestNeighbor produces a tour starting at 0 using greedy nearest neighbor.
// Defensive: if distances are "infinite" or missing it still picks an unused node.
func nearestNeighbor(graph [][]int64) []int {
	n := len(graph)
	if n == 0 {
		return nil
	}
	tour := make([]int, 0, n)
	used := make([]bool, n)
	cur := 0
	tour = append(tour, cur)
	used[cur] = true

	for len(tour) < n {
		best := -1
		bestCost := MaxInt64
		for v := 0; v < n; v++ {
			if used[v] {
				continue
			}
			w := graph[cur][v]
			// treat extremely large values as "unfavorable" but allow fallback
			if w < bestCost {
				bestCost = w
				best = v
			}
		}
		// fallback: pick first unused (defensive)
		if best == -1 {
			for v := 0; v < n; v++ {
				if !used[v] {
					best = v
					break
				}
			}
		}
		// final defensive guard: if something went wrong, break to avoid infinite loop
		if best == -1 || used[best] {
			break
		}
		tour = append(tour, best)
		used[best] = true
		cur = best
	}
	// final sanity: if we didn't visit all vertices, append missing ones
	if len(tour) < n {
		inTour := make([]bool, n)
		for _, v := range tour {
			inTour[v] = true
		}
		for v := 0; v < n; v++ {
			if !inTour[v] {
				tour = append(tour, v)
			}
		}
	}
	return tour
}

// twoOpt improves an existing tour by performing 2-opt swaps until no improvement.
// Keeps tour[0] as fixed start to match your BnB formatting.
// Safety: a max-iteration cap prevents pathological infinite loops.
func twoOpt(graph [][]int64, tour []int) []int {
	n := len(tour)
	if n <= 3 {
		return tour
	}

	const maxIterations = 500_000_000
	iter := 0
	improved := true

	for improved && iter < maxIterations {
		iter++
		improved = false
		// first-improvement: apply first beneficial swap then restart
		for i := 1; i < n-1; i++ {
			for j := n - 1; j > i; j-- {
				a, b := tour[i-1], tour[i]
				c, d := tour[j], tour[(j+1)%n]
				// delta = (a-c) + (b-d) - (a-b) - (c-d)
				// compute using int64 to match graph values
				delta := graph[a][c] + graph[b][d] - graph[a][b] - graph[c][d]
				if delta < 0 {
					for l, r := i, j; l < r; l, r = l+1, r-1 {
						tour[l], tour[r] = tour[r], tour[l]
					}
					improved = true
					break
				}
			}
			if improved {
				break
			}
		}
	}

	if iter >= maxIterations {
		fmt.Println("twoOpt: reached max iteration cap; stopping early")
	}
	return tour
}

// formatTourAsEdges converts a tour (vertex order) into the same formatted path array your BnB prints.
func formatTourAsEdges(graph [][]int64, tour []int) ([]string, int64) {
	path := []string{}
	n := len(tour)
	for i := 0; i < n-1; i++ {
		u, v := tour[i], tour[i+1]
		path = append(path, fmt.Sprintf("EDGE=[%d-%d] WEIGHT=[%d]", u, v, graph[u][v]))
	}
	// closing edge
	path = append(path, fmt.Sprintf("EDGE=[%d-%d] WEIGHT=[%d]", tour[n-1], tour[0], graph[tour[n-1]][tour[0]]))
	total := tourCost(graph, tour)
	return path, total
}

func mainASTAROptim(vertices int) {
	graph := generateGraph(vertices)

	startTime := time.Now()

	// 1. Nearest neighbor
	heurTour := nearestNeighbor(graph)
	// 2. Improve with 2-opt (local search)
	heurTour = twoOpt(graph, heurTour)
	path, cost := formatTourAsEdges(graph, heurTour)

	elapsed := time.Since(startTime)
	fmt.Println("Heuristic solution (Nearest-Neighbor + 2-opt):")
	for _, step := range path {
		fmt.Println(step)
	}
	fmt.Println("Cost:", cost)
	fmt.Println("Execution Time:", elapsed)
}
