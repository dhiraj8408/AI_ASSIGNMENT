package main

import (
	"fmt"
)

type Actions struct {
	ParentState string
	Action      string
}

var StateMapBDSForward = map[string]Actions{}
var StateMapBDSBackward = map[string]Actions{}
var NodesExpandedBDS int = 0

func SolveBDS(InitialConfiguration [][]rune, GoalConfiguration [][]rune, BlankIdxRowI int, BlankIdxColI int, BlankIdxRowG int, BlankIdxColG int, ParentI Actions, ParentG Actions) bool {
	KeyValI := GetHashKeyHelper(InitialConfiguration)
	KeyValG := GetHashKeyHelper(GoalConfiguration)

	// Check if states already visited
	if _, exists := StateMapBDSForward[KeyValI]; exists {
		return false
	}
	if _, exists := StateMapBDSBackward[KeyValG]; exists {
		return false
	}

	// Mark states as visited with correct parent tracking BEFORE intersection check
	StateMapBDSForward[KeyValI] = ParentI
	StateMapBDSBackward[KeyValG] = ParentG
	NodesExpandedBDS++

	// Check if the two searches have met (intersection found)
	if KeyValI == KeyValG {
		fmt.Printf("Direct intersection found: %s\n", KeyValI)
		return true
	}

	// Check if forward state exists in backward map or vice versa
	if _, exists := StateMapBDSBackward[KeyValI]; exists {
		fmt.Printf("Forward state found in backward map: %s\n", KeyValI)
		return true
	}
	if _, exists := StateMapBDSForward[KeyValG]; exists {
		fmt.Printf("Backward state found in forward map: %s\n", KeyValG)
		return true
	}

	directions := []struct {
		dRow, dCol int
		name       string
	}{
		{-1, 0, "UP"},
		{1, 0, "DOWN"},
		{0, -1, "LEFT"},
		{0, 1, "RIGHT"},
	}

	for _, dirI := range directions {
		NewBlankRowIdx := BlankIdxRowI + dirI.dRow
		NewBlankColIdx := BlankIdxColI + dirI.dCol

		for _, dirG := range directions {
			NewBlankRowIdxG := BlankIdxRowG + dirG.dRow
			NewBlankColIdxG := BlankIdxColG + dirG.dCol

			if NewBlankRowIdxG >= 0 && NewBlankRowIdxG < 3 && NewBlankColIdxG >= 0 && NewBlankColIdxG < 3 {
				if NewBlankRowIdx >= 0 && NewBlankRowIdx < 3 && NewBlankColIdx >= 0 && NewBlankColIdx < 3 {

					actionI := fmt.Sprintf("I:B -> %s -> %c", dirI.name, InitialConfiguration[NewBlankRowIdx][NewBlankColIdx])
					actionG := fmt.Sprintf("G:B -> %s -> %c", dirG.name, GoalConfiguration[NewBlankRowIdxG][NewBlankColIdxG])

					InitialConfiguration[BlankIdxRowI][BlankIdxColI], InitialConfiguration[NewBlankRowIdx][NewBlankColIdx] =
						InitialConfiguration[NewBlankRowIdx][NewBlankColIdx], InitialConfiguration[BlankIdxRowI][BlankIdxColI]

					GoalConfiguration[BlankIdxRowG][BlankIdxColG], GoalConfiguration[NewBlankRowIdxG][NewBlankColIdxG] =
						GoalConfiguration[NewBlankRowIdxG][NewBlankColIdxG], GoalConfiguration[BlankIdxRowG][BlankIdxColG]

					NewParentI := Actions{ParentState: KeyValI, Action: actionI}
					NewParentG := Actions{ParentState: KeyValG, Action: actionG}

					if SolveBDS(InitialConfiguration, GoalConfiguration, NewBlankRowIdx, NewBlankColIdx, NewBlankRowIdxG, NewBlankColIdxG, NewParentI, NewParentG) {
						return true
					}

					InitialConfiguration[BlankIdxRowI][BlankIdxColI], InitialConfiguration[NewBlankRowIdx][NewBlankColIdx] =
						InitialConfiguration[NewBlankRowIdx][NewBlankColIdx], InitialConfiguration[BlankIdxRowI][BlankIdxColI]

					GoalConfiguration[BlankIdxRowG][BlankIdxColG], GoalConfiguration[NewBlankRowIdxG][NewBlankColIdxG] =
						GoalConfiguration[NewBlankRowIdxG][NewBlankColIdxG], GoalConfiguration[BlankIdxRowG][BlankIdxColG]
				}
			}
		}
	}

	return false
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

	// Reconstruct forward path (from initial to meeting point)
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

	// Add initial state
	if current == initialKey {
		forwardPath = append([]string{initialKey}, forwardPath...)
	}

	// Reconstruct backward path (from meeting point to goal)
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

	// Add goal state if we reached it
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

	// Reset state maps and counter
	StateMapBDSForward = make(map[string]Actions)
	StateMapBDSBackward = make(map[string]Actions)
	NodesExpandedBDS = 0

	// Get initial and goal keys for reconstruction
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

	// Find goal blank position
	BlankIdxRowG, BlankIdxColG := 2, 2 // Goal state has blank at bottom-right

	fmt.Println("Starting Bidirectional DFS Search (BDS)...")

	if SolveBDS(InitialConfiguration, GoalConfiguration, BlankIdxRowI, BlankIdxColI, BlankIdxRowG, BlankIdxColG, ParentI, ParentG) {
		ReconstructPathFromStateMaps(initialKey, goalKey, &KeyValsPlayed, &ActionsPlayed)

		fmt.Printf("Solution found with BDS! Total nodes expanded: %d\n", NodesExpandedBDS)
		fmt.Printf("Total steps in solution: %d\n", len(KeyValsPlayed))
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
		fmt.Println("No solution found with BDS.")
	}
}
