package main

import (
	"container/heap"
	"fmt"
)

type StateNode struct {
	State        string
	ActionPlayed string
	F_N          int
	G_N          int
	Parent       *StateNode
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

var IndexMap = map[rune][2]int{
	'1': {0, 0}, '2': {0, 1}, '3': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'7': {2, 0}, '8': {2, 1}, 'B': {2, 2},
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDis(state string) int {
	distance := 0
	for i, c := range state {
		if c != 'B' {
			row := i / 3
			col := i % 3
			target := IndexMap[c]
			distance += absolute(row-target[0]) + absolute(col-target[1])
		}
	}
	return distance
}

func SolveASTAR(start, goal string) ([]string, []string, int, bool) {
	fringeList := &PriorityQueue{}
	heap.Init(fringeList)

	startNode := &StateNode{
		State:        start,
		ActionPlayed: "START",
		G_N:          0,
		F_N:          ManhattanDis(start),
		Parent:       nil,
	}
	heap.Push(fringeList, startNode)

	visited := make(map[string]int) // state → best g(n)

	directions := []struct {
		dRow, dCol int
		name       string
	}{
		{-1, 0, "UP"},
		{1, 0, "DOWN"},
		{0, -1, "LEFT"},
		{0, 1, "RIGHT"},
	}

	nodesExpanded := 0

	for fringeList.Len() > 0 {
		node := heap.Pop(fringeList).(*StateNode)
		state := node.State
		nodesExpanded++

		if state == goal {
			pathStates := []string{}
			pathActions := []string{}
			for n := node; n != nil; n = n.Parent {
				pathStates = append([]string{n.State}, pathStates...)
				if n.ActionPlayed != "START" {
					pathActions = append([]string{n.ActionPlayed}, pathActions...)
				}
			}
			return pathStates, pathActions, nodesExpanded, true
		}

		if gBest, ok := visited[state]; ok && gBest <= node.G_N {
			continue
		}
		visited[state] = node.G_N

		// locate blank
		blankIdx := 0
		for i, c := range state {
			if c == 'B' {
				blankIdx = i
				break
			}
		}
		rowIdx := blankIdx / 3
		colIdx := blankIdx % 3

		for _, dir := range directions {
			newRow, newCol := rowIdx+dir.dRow, colIdx+dir.dCol
			if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
				newBlankIdx := newRow*3 + newCol
				stateRunes := []rune(state)

				movedTile := stateRunes[newBlankIdx]

				stateRunes[blankIdx], stateRunes[newBlankIdx] =
					stateRunes[newBlankIdx], stateRunes[blankIdx]

				newState := string(stateRunes)
				g := node.G_N + 1
				h := ManhattanDis(newState)
				f := g + h

				child := &StateNode{
					State:        newState,
					ActionPlayed: fmt.Sprintf("%c %s", movedTile, dir.name),
					G_N:          g,
					F_N:          f,
					Parent:       node,
				}
				heap.Push(fringeList, child)
			}
		}
	}
	return nil, nil, nodesExpanded, false
}
func mainASTAR(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, uiEnabled bool) {
	// Convert board to string keys
	initialState := GetHashKeyHelper(InitialConfiguration)
	goalState := GetHashKeyHelper(GoalConfiguration)

	fmt.Println("Starting A* Search (ASTAR)...")

	// Run solver
	keyValsPlayed, actionsPlayed, nodesExpanded, success := SolveASTAR(initialState, goalState)

	if success {
		fmt.Printf("Solution found with ASTAR! Total nodes expanded: %d\n", nodesExpanded)
		fmt.Printf("Total steps in solution: %d\n", len(keyValsPlayed)-1) // exclude initial

		// Write output to file
		fileManager := NewFileOutputManager("output_ASTAR.txt")
		fmt.Println("Writing solution to output_ASTAR.txt...")

		for i, key := range keyValsPlayed {
			if i == 0 {
				fileManager.PrintBoardWithAction(key, "", i+1) // initial state has no action
			} else {
				fileManager.PrintBoardWithAction(key, actionsPlayed[i-1], i+1)
			}
		}

		// Generate visualization
		GenerateWebVisualization(keyValsPlayed, actionsPlayed, "ASTAR", nodesExpanded, uiEnabled)

	} else {
		fmt.Println("No solution found with ASTAR.")
	}
}
