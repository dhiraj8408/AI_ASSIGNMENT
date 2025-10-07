package main

import (
	"fmt"
	"time"
)

var StateMapDFS = map[string]int{}
var NodesExpandedDFS int = 0

func SolvePuzzleGetKeyValsDFS(InitialConfiguration [][]rune, GoalConfiguration [][]rune,
	BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string) bool {

	KeyVal := GetHashKeyHelper(InitialConfiguration)

	if _, ok := StateMapDFS[KeyVal]; ok {
		return false
	}

	StateMapDFS[KeyVal] = 1
	*KeyValsPlayed = append(*KeyValsPlayed, KeyVal)

	if CheckIfGoalStateReachedHelper(InitialConfiguration, GoalConfiguration) {
		return true
	}

	NodesExpandedDFS++

	directions := []struct {
		dr, dc int
		action string
	}{
		{-1, 0, "UP"},
		{+1, 0, "DOWN"},
		{0, -1, "LEFT"},
		{0, +1, "RIGHT"},
	}

	for _, dir := range directions {
		newR, newC := BlankIdxRow+dir.dr, BlankIdxCol+dir.dc
		if newR >= 0 && newR < 3 && newC >= 0 && newC < 3 {
			tileSwapped := InitialConfiguration[newR][newC]
			action := fmt.Sprintf("B -> %s -> %c", dir.action, tileSwapped)
			*ActionsPlayed = append(*ActionsPlayed, action)

			InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[newR][newC] =
				InitialConfiguration[newR][newC], InitialConfiguration[BlankIdxRow][BlankIdxCol]

			if SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, newR, newC, KeyValsPlayed, ActionsPlayed) {
				return true
			}

			InitialConfiguration[BlankIdxRow][BlankIdxCol], InitialConfiguration[newR][newC] =
				InitialConfiguration[newR][newC], InitialConfiguration[BlankIdxRow][BlankIdxCol]

			*ActionsPlayed = (*ActionsPlayed)[:len(*ActionsPlayed)-1]
		}
	}

	*KeyValsPlayed = (*KeyValsPlayed)[:len(*KeyValsPlayed)-1]

	return false
}

func mainDFS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, uiEnabled bool) {
	KeyValsPlayed := []string{}
	ActionsPlayed := []string{}
	fileManager := NewFileOutputManager("output_dfs.txt")

	fmt.Println("Starting Depth-First Search (DFS)...")
	startTime := time.Now()

	if SolvePuzzleGetKeyValsDFS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, &KeyValsPlayed, &ActionsPlayed) {
		endTime := time.Now()
		fmt.Printf("Solution found with DFS! Total nodes expanded: %d\n", NodesExpandedDFS)
		fmt.Printf("Total steps in solution: %d\n", len(KeyValsPlayed))
		fmt.Printf("Total execution time (DFS): %v\n", endTime.Sub(startTime))
		fmt.Println("Writing solution to output_dfs.txt...")

		for i, key := range KeyValsPlayed {
			if i > 0 {
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}

		GenerateWebVisualization(KeyValsPlayed, ActionsPlayed, "DFS", NodesExpandedDFS, uiEnabled)
	} else {
		endTime := time.Now()
		fmt.Println("No solution found with DFS.")
		fmt.Printf("Total execution time (DFS): %v\n", endTime.Sub(startTime))
	}
}
