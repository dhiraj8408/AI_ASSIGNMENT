package main

import (
	"fmt"
)

// Helper function to create a deep copy of the board
func CopyBoard(original [][]rune) [][]rune {
	copy := make([][]rune, len(original))
	for i := range original {
		copy[i] = make([]rune, len(original[i]))
		for j := range original[i] {
			copy[i][j] = original[i][j]
		}
	}
	return copy
}

func main() {
	fmt.Println("=== 8-PUZZLE SOLVER COMPARISON ===")

	uiEnabled := false

	fmt.Println("Generating test case...")
	InitialConfigurationGlobal, BlankIdxRowGlobal, BlankIdxColGlobal := GenerateTestCaseHelper()
	GoalConfiguration := GetGoalConfiguration()

	fmt.Printf("Initial Board Configuration:\n")
	PrintBoardHelper(InitialConfigurationGlobal)
	fmt.Printf("Blank position: Row %d, Col %d\n\n", BlankIdxRowGlobal, BlankIdxColGlobal)

	fmt.Println("Target Goal Configuration:")
	PrintBoardHelper(GoalConfiguration)
	fmt.Println()

	fmt.Println("1. Solving Board via DFS...")
	dfsBoard := CopyBoard(InitialConfigurationGlobal)
	mainDFS(dfsBoard, GoalConfiguration, BlankIdxRowGlobal, BlankIdxColGlobal, uiEnabled)

	fmt.Println("\n==================================================")

	fmt.Println("2. Solving Board via IDS...")
	idsBoard := CopyBoard(InitialConfigurationGlobal)
	mainIDS(idsBoard, GoalConfiguration, BlankIdxRowGlobal, BlankIdxColGlobal, uiEnabled)

	fmt.Println("\n==================================================")

	fmt.Println("3. Solving Board via BDS...")
	bdsBoard := CopyBoard(InitialConfigurationGlobal)
	mainBDS(bdsBoard, GoalConfiguration, BlankIdxRowGlobal, BlankIdxColGlobal, uiEnabled)

	fmt.Println("\n==================================================")

	fmt.Println("4. Solving Board via A*...")
	ASTARBoard := CopyBoard(InitialConfigurationGlobal)
	mainASTAR(ASTARBoard, GoalConfiguration, BlankIdxRowGlobal, BlankIdxColGlobal, uiEnabled)

	fmt.Println("\n==================================================")

	fmt.Println("\n=== COMPARISON COMPLETE ===")
	if uiEnabled {
		fmt.Println("\n Web visualizations have been generated!")
		fmt.Println("   Open the following HTML files in your browser:")
		fmt.Println("   - puzzle_solution_dfs.html")
		fmt.Println("   - puzzle_solution_ids.html")
		fmt.Println("   - puzzle_solution_bds.html")
	}
}
