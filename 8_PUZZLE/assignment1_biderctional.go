package main

import (
	"fmt"
	"time"
)

type Actions struct {
	ParentState string
	Action      string
}

var StateMapBDSForward = map[string]Actions{}
var StateMapBDSBackward = map[string]Actions{}
var NodesExpandedBDS int = 0

func CopyAndSwap(config [][]rune, r1, c1, r2, c2 int) [][]rune {
	newConfig := make([][]rune, 3)
	for i := 0; i < 3; i++ {
		newConfig[i] = make([]rune, 3)
		copy(newConfig[i], config[i])
	}
	newConfig[r1][c1], newConfig[r2][c2] = newConfig[r2][c2], newConfig[r1][c1]
	return newConfig
}

func SolveBDS_BFS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRowI int, BlankIdxColI int, BlankIdxRowG int, BlankIdxColG int, ParentI Actions, ParentG Actions) bool {
	startTime := time.Now()
	fmt.Printf("BDS BFS execution started at: %s\n", startTime.Format(time.RFC3339Nano))

	type Node struct {
		Config      [][]rune
		BlankRow    int
		BlankCol    int
		ParentState string
		Action      string
	}

	forwardQueue := []Node{{Config: InitialConfiguration, BlankRow: BlankIdxRowI, BlankCol: BlankIdxColI, ParentState: ParentI.ParentState, Action: ParentI.Action}}
	backwardQueue := []Node{{Config: GoalConfiguration, BlankRow: BlankIdxRowG, BlankCol: BlankIdxColG, ParentState: ParentG.ParentState, Action: ParentG.Action}}

	StateMapBDSForward = map[string]Actions{}
	StateMapBDSBackward = map[string]Actions{}
	NodesExpandedBDS = 0

	directions := []struct {
		dRow, dCol int
		name       string
	}{
		{-1, 0, "UP"},
		{1, 0, "DOWN"},
		{0, -1, "LEFT"},
		{0, 1, "RIGHT"},
	}

	for len(forwardQueue) > 0 || len(backwardQueue) > 0 {
		if len(forwardQueue) > 0 {
			curr := forwardQueue[0]
			forwardQueue = forwardQueue[1:]
			key := GetHashKeyHelper(curr.Config)
			if _, exists := StateMapBDSForward[key]; exists {
				continue
			}
			StateMapBDSForward[key] = Actions{ParentState: curr.ParentState, Action: curr.Action}
			NodesExpandedBDS++

			if _, exists := StateMapBDSBackward[key]; exists {
				fmt.Printf("Intersection found at: %s\n", key)
				return true
			}

			for _, dir := range directions {
				newRow, newCol := curr.BlankRow+dir.dRow, curr.BlankCol+dir.dCol
				if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
					newConfig := CopyAndSwap(curr.Config, curr.BlankRow, curr.BlankCol, newRow, newCol)
					action := fmt.Sprintf("I:B -> %s -> %c", dir.name, curr.Config[newRow][newCol])
					forwardQueue = append(forwardQueue, Node{
						Config:      newConfig,
						BlankRow:    newRow,
						BlankCol:    newCol,
						ParentState: key,
						Action:      action,
					})
				}
			}
		}

		if len(backwardQueue) > 0 {
			curr := backwardQueue[0]
			backwardQueue = backwardQueue[1:]
			key := GetHashKeyHelper(curr.Config)
			if _, exists := StateMapBDSBackward[key]; exists {
				continue
			}
			StateMapBDSBackward[key] = Actions{ParentState: curr.ParentState, Action: curr.Action}
			NodesExpandedBDS++

			if _, exists := StateMapBDSForward[key]; exists {
				fmt.Printf("Intersection found at: %s\n", key)
				return true
			}

			for _, dir := range directions {
				newRow, newCol := curr.BlankRow+dir.dRow, curr.BlankCol+dir.dCol
				if newRow >= 0 && newRow < 3 && newCol >= 0 && newCol < 3 {
					newConfig := CopyAndSwap(curr.Config, curr.BlankRow, curr.BlankCol, newRow, newCol)
					action := fmt.Sprintf("G:B -> %s -> %c", dir.name, curr.Config[newRow][newCol])
					backwardQueue = append(backwardQueue, Node{
						Config:      newConfig,
						BlankRow:    newRow,
						BlankCol:    newCol,
						ParentState: key,
						Action:      action,
					})
				}
			}
		}
	}
	return false
}

func FlattenAndSwap(config [][]rune, r1, c1, r2, c2 int) [][]rune {
	flat := make([]rune, 9)
	idx := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			flat[idx] = config[i][j]
			idx++
		}
	}
	pos1 := r1*3 + c1
	pos2 := r2*3 + c2
	flat[pos1], flat[pos2] = flat[pos2], flat[pos1]
	// Convert back to [][]rune
	newConfig := make([][]rune, 3)
	for i := 0; i < 3; i++ {
		newConfig[i] = make([]rune, 3)
		for j := 0; j < 3; j++ {
			newConfig[i][j] = flat[i*3+j]
		}
	}
	return newConfig
}

func ReconstructPathFromStateMaps(initialKey string, goalKey string, KeyValsPlayed *[]string, ActionsPlayed *[]string) {
	var meetingPoint string
	var foundIntersection bool

	for forwardKey := range StateMapBDSForward {
		if _, exists := StateMapBDSBackward[forwardKey]; exists {
			meetingPoint = forwardKey
			foundIntersection = true
			break
		}
	}

	if !foundIntersection {
		for backwardKey := range StateMapBDSBackward {
			if _, exists := StateMapBDSForward[backwardKey]; exists {
				meetingPoint = backwardKey
				foundIntersection = true
				break
			}
		}
	}

	if !foundIntersection {
		fmt.Println("No intersection found in state maps!")
		return
	}

	fmt.Printf("Meeting point found: %s\n", meetingPoint)

	forwardPath := []string{}
	forwardActions := []string{}
	current := meetingPoint

	for current != "" && current != initialKey {
		forwardPath = append([]string{current}, forwardPath...)
		if action, exists := StateMapBDSForward[current]; exists {
			if action.Action != "START" && action.Action != "" {
				forwardActions = append([]string{action.Action}, forwardActions...)
			}
			current = action.ParentState
		} else {
			break
		}
	}

	if current == initialKey {
		forwardPath = append([]string{initialKey}, forwardPath...)
	}

	backwardPath := []string{}
	backwardActions := []string{}
	current = meetingPoint

	for current != "" && current != goalKey {
		if action, exists := StateMapBDSBackward[current]; exists {
			if action.ParentState != "" && action.ParentState != goalKey {
				backwardPath = append(backwardPath, action.ParentState)
				if action.Action != "START" && action.Action != "" {
					backwardActions = append(backwardActions, action.Action)
				}
			}
			current = action.ParentState
		} else {
			break
		}
	}

	if current == goalKey {
		backwardPath = append(backwardPath, goalKey)
	}

	fmt.Printf("Forward path length: %d\n", len(forwardPath))
	fmt.Printf("Backward path length: %d\n", len(backwardPath))

	*KeyValsPlayed = forwardPath
	if len(backwardPath) > 0 {
		*KeyValsPlayed = append(*KeyValsPlayed, backwardPath...)
	}

	*ActionsPlayed = forwardActions
	*ActionsPlayed = append(*ActionsPlayed, backwardActions...)
}

func mainBDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRowI int, BlankIdxColI int, uiEnabled bool) {
	KeyValsPlayed := []string{}
	ActionsPlayed := []string{}

	StateMapBDSForward = make(map[string]Actions)
	StateMapBDSBackward = make(map[string]Actions)
	NodesExpandedBDS = 0

	initialKey := GetHashKeyHelper(InitialConfiguration)
	goalKey := GetHashKeyHelper(GetGoalConfiguration())

	ParentI := Actions{
		ParentState: "",
		Action:      "START",
	}
	ParentG := Actions{
		ParentState: "",
		Action:      "START",
	}

	fileManager := NewFileOutputManager("output_BDS.txt")

	BlankIdxRowG, BlankIdxColG := 2, 2

	fmt.Println("Starting Bidirectional BFS Search (BDS)...")
	startTime := time.Now()

	if SolveBDS_BFS(InitialConfiguration, GoalConfiguration, BlankIdxRowI, BlankIdxColI, BlankIdxRowG, BlankIdxColG, ParentI, ParentG) {
		endTime := time.Now()
		ReconstructPathFromStateMaps(initialKey, goalKey, &KeyValsPlayed, &ActionsPlayed)

		fmt.Printf("Solution found with BDS! Total nodes expanded: %d\n", NodesExpandedBDS)
		fmt.Printf("Total steps in solution: %d\n", len(ActionsPlayed))
		fmt.Printf("Total execution time (BDS): %v\n", endTime.Sub(startTime))
		fmt.Println("Writing solution to output_BDS.txt...")

		for i, key := range KeyValsPlayed {
			if i > 0 && i-1 < len(ActionsPlayed) {
				fileManager.PrintBoardWithAction(key, ActionsPlayed[i-1], i+1)
			} else {
				fileManager.PrintBoardWithAction(key, "", i+1)
			}
		}

		GenerateWebVisualization(KeyValsPlayed, ActionsPlayed, "BDS", NodesExpandedBDS, uiEnabled)
	} else {
		endTime := time.Now()
		fmt.Println("No solution found with BDS.")
		fmt.Printf("Total execution time (BDS): %v\n", endTime.Sub(startTime))
	}
}
