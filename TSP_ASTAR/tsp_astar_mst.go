package main

import (
	"container/heap"
	"fmt"
	"time"
)

func heuristic_MST(currentVertex int, unvisitedVertices *Set[int], graph [][]int64) int64 {
	if unvisitedVertices.Size() == 0 {
		return graph[currentVertex][0]
	}
	vertices := unvisitedVertices.Values()
	n := len(vertices)
	if n == 1 {
		return graph[currentVertex][vertices[0]] + graph[vertices[0]][0]
	}

	mstCost := int64(0)
	inMST := make(map[int]bool)
	start := vertices[0]
	inMST[start] = true

	pq := &edgeHeap{}
	heap.Init(pq)
	for _, v := range vertices {
		if v != start {
			heap.Push(pq, Edge{cost: graph[start][v], to: v})
		}
	}

	for pq.Len() > 0 && len(inMST) < n {
		edge := heap.Pop(pq).(Edge)
		if inMST[edge.to] {
			continue
		}
		inMST[edge.to] = true
		mstCost += edge.cost

		for _, v := range vertices {
			if !inMST[v] {
				heap.Push(pq, Edge{cost: graph[edge.to][v], to: v})
			}
		}
	}

	return mstCost
}

func tsp_helper_astar(graph [][]int64, vertices int) ([]string, int64, int, bool) {
	fringeList := &PriorityQueue{}
	heap.Init(fringeList)

	unvisited := NewSet[int]()
	for i := 1; i < vertices; i++ {
		unvisited.Add(i)
	}

	startNode := &StateNode{
		State:        GetStateKey(unvisited),
		ActionPlayed: "START",
		G_N:          0,
		F_N:          heuristic_MST(0, unvisited, graph),
		Parent:       nil,
		Vertex:       0,
		Unvisited:    unvisited,
	}

	heap.Push(fringeList, startNode)
	nodesExpanded := 0
	visitedStates := make(map[string]int64)

	for fringeList.Len() > 0 {
		node := heap.Pop(fringeList).(*StateNode)
		nodesExpanded++

		if node.Unvisited.Size() == 0 {
			totalCost := node.G_N + graph[node.Vertex][0]

			pathActions := []string{}
			for n := node; n != nil; n = n.Parent {
				if n.ActionPlayed != "START" {
					pathActions = append([]string{n.ActionPlayed}, pathActions...)
				}
			}

			pathActions = append(pathActions, fmt.Sprintf("EDGE=[%d-%d] WEIGHT=[%d]", node.Vertex, 0, graph[node.Vertex][0]))

			return pathActions, totalCost, nodesExpanded, true
		}

		if gBest, ok := visitedStates[node.State]; ok && gBest <= node.G_N {
			continue
		}
		visitedStates[node.State] = node.G_N

		for _, nextVertex := range node.Unvisited.Values() {
			newUnvisited := node.Unvisited.Copy()
			newUnvisited.Remove(nextVertex)

			g_N := node.G_N + graph[node.Vertex][nextVertex]
			h_N := heuristic_MST(nextVertex, newUnvisited, graph)
			f_N := g_N + h_N

			childNode := &StateNode{
				State:        GetStateKey(newUnvisited),
				ActionPlayed: fmt.Sprintf("EDGE=[%d-%d] WEIGHT=[%d]", node.Vertex, nextVertex, graph[node.Vertex][nextVertex]),
				G_N:          g_N,
				F_N:          f_N,
				Parent:       node,
				Vertex:       nextVertex,
				Unvisited:    newUnvisited,
			}
			heap.Push(fringeList, childNode)
		}
	}
	return nil, 0, nodesExpanded, false
}

func mainASTAR(vertices int) {
	graph := generateGraph(vertices)
	startTime := time.Now()
	actionsPlayed, cost, nodesExpanded, success := tsp_helper_astar(graph, vertices)
	if success {
		fmt.Println("Solution found:")
		for _, action := range actionsPlayed {
			fmt.Println(action)
		}
		fmt.Println("Cost:", cost)
		fmt.Println("Nodes Expanded:", nodesExpanded)
		endTime := time.Now()
		fmt.Printf("TSP[ASTAR H(N) = MST] Execution Time: %v\n", endTime.Sub(startTime))
	} else {
		fmt.Println("No solution found.")
		endTime := time.Now()
		fmt.Printf("TSP[ASTAR H(N) = MST] Execution Time: %v\n", endTime.Sub(startTime))
	}
}
