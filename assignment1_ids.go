package main

import (
	"fmt"
)

var NodesExpandedIDS int = 0

func SolvePuzzleGetKeyValsIDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string) bool {
	Solved := false
	MaxDepth := 50
	NodesExpandedIDS = 0 // Reset counter for IDS

	for depth := 0; depth < MaxDepth; depth++ {
		fmt.Printf("Searching at depth %d...\n", depth)

		// Clear the states before restarting the iteration
		var StateMapIDS = map[string]int{}

		// Clear previous attempts
		*KeyValsPlayed = []string{}
		*ActionsPlayed = []string{}

		Solved = SolveIDSHelper(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, KeyValsPlayed, ActionsPlayed, 0, depth, StateMapIDS)
		if Solved {
			return true
		}
	}

	return false
}

func SolveIDSHelper(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string, CurrentDepth int, MaxDepth int, StateMapIDS map[string]int) bool {
	KeyVal := GetHashKeyHelper(InitialConfiguration)

	// Check if we've seen this state at this depth or shallower
	if depth, exists := StateMapIDS[KeyVal]; exists && depth <= CurrentDepth {
		return false
	}

	StateMapIDS[KeyVal] = CurrentDepth

	// Check if goal state is reached
	if CheckIfGoalStateReachedHelper(InitialConfiguration, GoalConfiguration) {
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)
		return true
	}

	// If we've reached maximum depth, don't expand further
	if CurrentDepth >= MaxDepth {
		return false
	}

	// This is where we actually "expand" a node - we're going to try all possible moves
	*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)
	NodesExpandedIDS++ // Count this as an expanded node

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
	if SolvePuzzleGetKeyValsIDS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, &KeyValsPlayed, &ActionsPlayed) {
		fmt.Printf("Solution found with IDS! Total nodes expanded: %d\n", NodesExpandedIDS)
		fmt.Printf("Total steps in solution: %d\n", len(KeyValsPlayed))
		fmt.Println("Writing solution to output_IDS.txt...")

		for i, key := range KeyValsPlayed {
			if i > 0 {
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}

		// Generate web visualization if UI is enabled
		GenerateWebVisualization(KeyValsPlayed, ActionsPlayed, "IDS", NodesExpandedIDS, uiEnabled)
	} else {
		fmt.Println("No solution found with IDS.")
	}
}
