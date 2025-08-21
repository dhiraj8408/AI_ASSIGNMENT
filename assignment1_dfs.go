package main

import (
	"fmt"
)

var StateMapDFS = map[string]int{}
var NodesExpandedDFS int = 0

func SolvePuzzleGetKeyValsDFS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string) bool {
	KeyVal := GetHashKeyHelper(InitialConfiguration)
	if _, ok := StateMapDFS[KeyVal]; ok {
		return false
	}

	if CheckIfGoalStateReachedHelper(InitialConfiguration, GoalConfiguration) {
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)
		return true
	}

	Top, Bottom, Left, Right := false, false, false, false
	NodesExpandedDFS++

	if BlankIdxRow > 0 {
		StateMapDFS[KeyVal] = 1
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)

		// Record the action: B moves UP and swaps with the tile above
		tileSwapped := InitialConfiguration[BlankIdxRow-1][BlankIdxCol]
		action := fmt.Sprintf("B -> UP -> %c", tileSwapped)
		*ActionsPlayed = append(*ActionsPlayed, action)

		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow-1][BlankIdxCol] =
			InitialConfiguration[BlankIdxRow-1][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol]

		Top = SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow-1, BlankIdxCol, KeyValsPlayed, ActionsPlayed)
		if Top {
			return true
		}
		*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]
		*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow-1][BlankIdxCol] =
			InitialConfiguration[BlankIdxRow-1][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol]
	}

	if BlankIdxRow < 2 {
		StateMapDFS[KeyVal] = 1
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)

		// Record the action: B moves DOWN and swaps with the tile below
		tileSwapped := InitialConfiguration[BlankIdxRow+1][BlankIdxCol]
		action := fmt.Sprintf("B -> DOWN -> %c", tileSwapped)
		*ActionsPlayed = append(*ActionsPlayed, action)

		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow+1][BlankIdxCol] =
			InitialConfiguration[BlankIdxRow+1][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol]

		Bottom = SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow+1, BlankIdxCol, KeyValsPlayed, ActionsPlayed)
		if Bottom {
			return true
		}
		*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]
		*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow+1][BlankIdxCol] =
			InitialConfiguration[BlankIdxRow+1][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol]

	}

	if BlankIdxCol > 0 {
		StateMapDFS[KeyVal] = 1
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)

		// Record the action: B moves LEFT and swaps with the tile on the left
		tileSwapped := InitialConfiguration[BlankIdxRow][BlankIdxCol-1]
		action := fmt.Sprintf("B -> LEFT -> %c", tileSwapped)
		*ActionsPlayed = append(*ActionsPlayed, action)

		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol-1] =
			InitialConfiguration[BlankIdxRow][BlankIdxCol-1], InitialConfiguration[BlankIdxRow][BlankIdxCol]

		Left = SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol-1, KeyValsPlayed, ActionsPlayed)
		if Left {
			return true
		}
		*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]
		*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol-1] =
			InitialConfiguration[BlankIdxRow][BlankIdxCol-1], InitialConfiguration[BlankIdxRow][BlankIdxCol]
	}

	if BlankIdxCol < 2 {
		StateMapDFS[KeyVal] = 1
		*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)

		// Record the action: B moves RIGHT and swaps with the tile on the right
		tileSwapped := InitialConfiguration[BlankIdxRow][BlankIdxCol+1]
		action := fmt.Sprintf("B -> RIGHT -> %c", tileSwapped)
		*ActionsPlayed = append(*ActionsPlayed, action)

		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol+1] =
			InitialConfiguration[BlankIdxRow][BlankIdxCol+1], InitialConfiguration[BlankIdxRow][BlankIdxCol]

		Right = SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol+1, KeyValsPlayed, ActionsPlayed)
		if Right {
			return true
		}
		*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]
		*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[BlankIdxRow][BlankIdxCol+1] =
			InitialConfiguration[BlankIdxRow][BlankIdxCol+1], InitialConfiguration[BlankIdxRow][BlankIdxCol]
	}

	return false
}

func mainDFS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, uiEnabled bool) {
	KeyValsPlayed := []string{}
	ActionsPlayed := []string{}
	fileManager := NewFileOutputManager("output_dfs.txt")

	fmt.Println("Starting Depth-First Search (DFS)...")
	if SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, &KeyValsPlayed, &ActionsPlayed) {
		fmt.Printf("Solution found with DFS! Total nodes expanded: %d\n", NodesExpandedDFS)
		fmt.Printf("Total steps in solution: %d\n", len(KeyValsPlayed))
		fmt.Println("Writing solution to output_dfs.txt...")

		for i, key := range KeyValsPlayed {
			if i > 0 {
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}

		// Generate web visualization if UI is enabled
		GenerateWebVisualization(KeyValsPlayed, ActionsPlayed, "DFS", NodesExpandedDFS, uiEnabled)
	} else {
		fmt.Println("No solution found with DFS.")
	}
}
