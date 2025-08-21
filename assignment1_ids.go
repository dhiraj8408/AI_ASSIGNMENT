earpackage main

import (
	"fmt"
)

func SolvePuzzleGetKeyValsIDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string, Depth int) bool {
	Solved := false
	MaxDepth := 50

	for depth := 0; depth < MaxDepth; depth++ {
		// Clear the states before restarting the iteration
		var StateMapIDS = map[string]int{}

		SolveIDSHelper(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, KeyValsPlayed, ActionsPlayed, 0, depth)
	}
}

func SolveIDSHelper(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRow int, BlankIdxCol int, KeyValsPlayed *[]string, ActionsPlayed *[]string, CurrentDepth int, MaxDepth int){
	
}

func main() {
	InitialConfiguration, BlankIdxRow, BlankIdxCol := GenerateTestCaseHelper()
	GoalConfiguration := GetGoalConfiguration()

	KeyValsPlayed := []string{}
	ActionsPlayed := []string{}
	fileManager := NewFileOutputManager("output_dfs.txt")

	if SolvePuzzleGetKeyValsIDS(InitialConfiguration, GoalConfiguration, BlankIdxRow, BlankIdxCol, &KeyValsPlayed, &ActionsPlayed, fileManager) {
		fmt.Println("Solution found!")
		for i, key := range KeyValsPlayed {
			fmt.Printf("Step %d:\n", i+1)
			if i > 0 {
				fmt.Printf("Action: %s\n", ActionsPlayed[i-1])
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fmt.Println("Initial state:")
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}
	} else {
		fmt.Println("No solution found.")
	}
}
