package main

import "fmt"

func main() {
	var vertices int
	fmt.Print("Enter number of vertices: ")
	fmt.Scanln(&vertices)

	if vertices <= 20 {
		mainASTAR(vertices)
	} else {
		mainASTAROptim(vertices)
	}
}
