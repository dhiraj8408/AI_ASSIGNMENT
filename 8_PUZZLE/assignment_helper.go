package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// GetHashKeyHelper generates a unique string key from a 2D matrix
func GetHashKeyHelper(Matrix [][]rune) string {
	Key := ""
	for _, row := range Matrix {
		for _, ch := range row {
			Key += string(ch)
		}
	}
	return Key
}

// CheckIfGoalStateReachedHelper compares two 2D matrices to see if they are identical
func CheckIfGoalStateReachedHelper(a [][]rune, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// GenerateRandomBoardHelper creates a randomly shuffled 3x3 puzzle board
func GenerateRandomBoardHelper() [][]rune {
	tiles := []rune{'1', '2', '3', '4', '5', '6', '7', '8', 'B'}

	rand.Seed(time.Now().UnixNano())

	shuffleCount := 10 + rand.Intn(20)
	for k := 0; k < shuffleCount; k++ {
		for i := len(tiles) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			tiles[i], tiles[j] = tiles[j], tiles[i]
		}
	}

	board := make([][]rune, 3)
	for i := range board {
		board[i] = make([]rune, 3)
		for j := range board[i] {
			board[i][j] = tiles[i*3+j]
		}
	}

	return board
}

// CheckIsSolvableHelper determines if a puzzle configuration is solvable
func CheckIsSolvableHelper(board [][]rune) bool {
	inversions := 0
	flat := []rune{}
	for _, row := range board {
		for _, val := range row {
			if val != 'B' {
				flat = append(flat, val)
			}
		}
	}
	for i := 0; i < len(flat); i++ {
		for j := i + 1; j < len(flat); j++ {
			if flat[i] > flat[j] {
				inversions++
			}
		}
	}
	return inversions%2 == 0
}

// PrintBoardHelper prints the board to console
func PrintBoardHelper(board [][]rune) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%c ", val)
		}
		fmt.Println()
	}
}

// GenerateTestCaseHelper generates a solvable random board and returns it with blank position
func GenerateTestCaseHelper() ([][]rune, int, int) {
	board := [][]rune{
		{'1', '3', '5'},
		{'B', '2', '6'},
		{'4', '7', '8'},
	}
	for !CheckIsSolvableHelper(board) {
		fmt.Println("Generated board is not solvable, generating a new one...")
		board = GenerateRandomBoardHelper()
	}
	fmt.Println("Generated solvable board:")
	PrintBoardHelper(board)
	var BlankIdxRow, BlankIdxCol int
	found := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'B' {
				BlankIdxRow, BlankIdxCol = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	return board, BlankIdxRow, BlankIdxCol
}

// FileOutputManager handles file output operations
type FileOutputManager struct {
	isFirstCall bool
	filename    string
}

// NewFileOutputManager creates a new file output manager
func NewFileOutputManager(filename string) *FileOutputManager {
	return &FileOutputManager{
		isFirstCall: true,
		filename:    filename,
	}
}

// PrintBoardFromKey writes a board state to file from its string representation
func (fom *FileOutputManager) PrintBoardFromKey(str string) {
	var file *os.File
	var err error

	if fom.isFirstCall {
		file, err = os.OpenFile(fom.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		fom.isFirstCall = false
	} else {
		file, err = os.OpenFile(fom.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 9; i++ {
		fmt.Fprintf(file, "%c ", str[i])
		if (i+1)%3 == 0 {
			fmt.Fprintln(file)
		}
	}
	fmt.Fprintln(file, "\n")
}

// PrintBoardWithAction writes a board state with action information to file
func (fom *FileOutputManager) PrintBoardWithAction(str string, action string, stepNum int) {
	var file *os.File
	var err error

	if fom.isFirstCall {
		file, err = os.OpenFile(fom.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		fom.isFirstCall = false
	} else {
		file, err = os.OpenFile(fom.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "Step %d:\n", stepNum)
	if action != "" {
		fmt.Fprintf(file, "Action: %s\n", action)
	}

	for i := 0; i < 9; i++ {
		fmt.Fprintf(file, "%c ", str[i])
		if (i+1)%3 == 0 {
			fmt.Fprintln(file)
		}
	}
	fmt.Fprintln(file, "\n")
}

// GetGoalConfiguration returns the standard goal configuration
func GetGoalConfiguration() [][]rune {
	return [][]rune{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', 'B'},
	}
}
