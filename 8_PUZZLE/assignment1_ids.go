package main

import (
	"fmt"
	"time"
)

var NodesExpandedIDS int = 0

func SolvePuzzleGetKeyValsIDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string) bool {
	Solved := false
	MaxDepth := 50
	NodesExpandedIDS = 0

	startTime := time.Now()
	fmt.Printf("IDS execution started at: %s\n", startTime.Format(time.RFC3339Nano))

	for depth := 0; depth < MaxDepth; depth++ {
		iterStart := time.Now()
		fmt.Printf("Searching at depth %d...\n", depth)

		var StateMapIDS = map[string]int{}

		*KeyValsPlayed = []string{}
		*ActionsPlayed = []string{}

		Solved = SolveIDSHelper(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, KeyValsPlayed, ActionsPlayed, 0, depth, StateMapIDS)
		iterEnd := time.Now()
		fmt.Printf("Iteration at depth %d took: %v\n", depth, iterEnd.Sub(iterStart))
		if Solved {
			totalEnd := time.Now()
			fmt.Printf("IDS completed at: %s\n", totalEnd.Format(time.RFC3339Nano))
			fmt.Printf("Total execution time: %v\n", totalEnd.Sub(startTime))
			return true
		}
	}

	totalEnd := time.Now()
	fmt.Printf("IDS completed at: %s\n", totalEnd.Format(time.RFC3339Nano))
	fmt.Printf("Total execution time (IDS): %v\n", totalEnd.Sub(startTime))
	return false
}

func SolveIDSHelper(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string, CurrentDepth int, MaxDepth int, StateMapIDS map[string]int) bool {
	KeyVal := GetHashKeyHelper(InitialConfiguration)

	if depth, exists := StateMapIDS[KeyVal]; exists && depth <= CurrentDepth {
		return false
	}

	StateMapIDS[KeyVal] = CurrentDepth

	if CheckIfGoalStateReachedHelper(InitialConfiguration, GoalConfiguration) {
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)
		return true
	}

	if CurrentDepth >= MaxDepth {
		return false
	}

	*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)
	NodesExpandedIDS++

	directions := []struct {
		dRow, dCol int
		name       string
	}{
		{-1, 0, "UP"},
		{1, 0, "DOWN"},
		{0, -1, "LEFT"},
		{0, 1, "RIGHT"},
	}

	for _, dir := range directions {
		NewBlankRowIdx := BlankIdxRow + dir.dRow
		NewBlankColIdx := BlankIdxCol + dir.dCol

		if NewBlankRowIdx >= 0 && NewBlankRowIdx < 3 && NewBlankColIdx >= 0 && NewBlankColIdx < 3 {
			action := fmt.Sprintf("B -> %s -> %c", dir.name, InitialConfiguration[NewBlankRowIdx][NewBlankColIdx])
			*ActionsPlayed = append(*ActionsPlayed, action)

			InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[NewBlankRowIdx][NewBlankColIdx] = InitialConfiguration[NewBlankRowIdx][NewBlankColIdx], InitialConfiguration[BlankIdxRow][BlankIdxCol]

			if SolveIDSHelper(InitialConfiguration, GoalConfiguration, NewBlankRowIdx, NewBlankColIdx, KeyValsPlayed, ActionsPlayed, CurrentDepth+1, MaxDepth, StateMapIDS) {
				return true
			}

			InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[NewBlankRowIdx][NewBlankColIdx] = InitialConfiguration[NewBlankRowIdx][NewBlankColIdx], InitialConfiguration[BlankIdxRow][BlankIdxCol]
			*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		}
	}
	*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]
	return false
}

func mainIDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, uiEnabled bool) {
	KeyValsPlayed := []string{}
	ActionsPlayed := []string{}
	fileManager := NewFileOutputManager("output_IDS.txt")

	fmt.Println("Starting Iterative Deepening Search (IDS)...")
	startTime := time.Now()

	if SolvePuzzleGetKeyValsIDS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, &KeyValsPlayed, &ActionsPlayed) {
		endTime := time.Now()
		fmt.Printf("Solution found with IDS! Total nodes expanded: %d\n", NodesExpandedIDS)
		fmt.Printf("Total steps in solution: %d\n", len(ActionsPlayed))
		fmt.Printf("Total execution time (IDS): %v\n", endTime.Sub(startTime))
		fmt.Println("Writing solution to output_IDS.txt...")

		for i, key := range KeyValsPlayed {
			if i > 0 {
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}

		GenerateWebVisualization(KeyValsPlayed, ActionsPlayed, "IDS", NodesExpandedIDS, uiEnabled)
	} else {
		endTime := time.Now()
		fmt.Println("No solution found with IDS.")
		fmt.Printf("Total execution time (IDS): %v\n", endTime.Sub(startTime))
	}
}
