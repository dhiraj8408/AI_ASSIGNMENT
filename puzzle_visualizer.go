package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// PuzzleVisualizer handles visual display of the puzzle solving process
type PuzzleVisualizer struct {
	delay time.Duration
}

// NewPuzzleVisualizer creates a new visualizer with specified delay between steps
func NewPuzzleVisualizer(delayMs int) *PuzzleVisualizer {
	return &PuzzleVisualizer{
		delay: time.Duration(delayMs) * time.Millisecond,
	}
}

// ClearScreen clears the terminal screen
func (pv *PuzzleVisualizer) ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// PrintBoardAnimated prints a board with fancy formatting and animations
func (pv *PuzzleVisualizer) PrintBoardAnimated(boardStr string, stepNum int, action string) {
	pv.ClearScreen()

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Printf("║           8-PUZZLE SOLVER            ║\n")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Printf("║ Step: %-3d                          ║\n", stepNum)
	if action != "" {
		fmt.Printf("║ Action: %-28s ║\n", action)
	} else {
		fmt.Printf("║ Action: %-28s ║\n", "Initial State")
	}
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║                                      ║")

	// Print the board with fancy borders
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Print("║          ┌───┬───┬───┐              ║\n")
			fmt.Print("║          │")
		}

		if boardStr[i] == 'B' {
			fmt.Print("   │")
		} else {
			fmt.Printf(" %c │", boardStr[i])
		}

		if (i+1)%3 == 0 {
			fmt.Print("              ║\n")
			if i < 6 {
				fmt.Print("║          ├───┼───┼───┤              ║\n")
			} else {
				fmt.Print("║          └───┴───┴───┘              ║\n")
			}
		}
	}

	fmt.Println("║                                      ║")
	fmt.Println("╚══════════════════════════════════════╝")

	// Add a progress indicator
	progress := float64(stepNum) / 20.0 // Assume max 20 steps for demo
	if progress > 1.0 {
		progress = 1.0
	}
	progressBar := ""
	for i := 0; i < 30; i++ {
		if float64(i) < progress*30 {
			progressBar += "█"
		} else {
			progressBar += "░"
		}
	}
	fmt.Printf("Progress: [%s] %.1f%%\n", progressBar, progress*100)

	time.Sleep(pv.delay)
}

// AnimateSolution animates the entire solution process
func (pv *PuzzleVisualizer) AnimateSolution(keyVals []string, actions []string) {
	fmt.Println("🎯 Starting 8-Puzzle Solution Animation...")
	fmt.Println("Press Ctrl+C to stop the animation")
	time.Sleep(2 * time.Second)

	for i, key := range keyVals {
		action := ""
		if i > 0 && i-1 < len(actions) {
			action = actions[i-1]
		}
		pv.PrintBoardAnimated(key, i+1, action)
	}

	// Final celebration
	pv.ClearScreen()
	fmt.Println("🎉 PUZZLE SOLVED! 🎉")
	fmt.Printf("Solution found in %d steps!\n", len(keyVals))
	fmt.Println("\n" + keyVals[len(keyVals)-1])

	// Display final board one more time
	finalBoard := keyVals[len(keyVals)-1]
	fmt.Println("\n╔═══════════════════════════════════════╗")
	fmt.Println("║           FINAL SOLUTION              ║")
	fmt.Println("╠═══════════════════════════════════════╣")
	fmt.Println("║                                       ║")

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Print("║          ┌───┬───┬───┐               ║\n")
			fmt.Print("║          │")
		}

		if finalBoard[i] == 'B' {
			fmt.Print("   │")
		} else {
			fmt.Printf(" %c │", finalBoard[i])
		}

		if (i+1)%3 == 0 {
			fmt.Print("               ║\n")
			if i < 6 {
				fmt.Print("║          ├───┼───┼───┤               ║\n")
			} else {
				fmt.Print("║          └───┴───┴───┘               ║\n")
			}
		}
	}

	fmt.Println("║                                       ║")
	fmt.Println("╚═══════════════════════════════════════╝")
}

// WebVisualizerHTML generates HTML file for web-based visualization
func (pv *PuzzleVisualizer) WebVisualizerHTML(keyVals []string, actions []string, filename string) {
	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>8-Puzzle Solver Visualization</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            margin: 0;
            padding: 20px;
            color: white;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            text-align: center;
        }
        .puzzle-board {
            display: inline-grid;
            grid-template-columns: repeat(3, 80px);
            grid-gap: 5px;
            background: #333;
            padding: 20px;
            border-radius: 15px;
            margin: 20px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.3);
        }
        .tile {
            width: 80px;
            height: 80px;
            background: linear-gradient(145deg, #f0f0f0, #d0d0d0);
            border-radius: 10px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 24px;
            font-weight: bold;
            color: #333;
            transition: all 0.3s ease;
            box-shadow: 0 4px 8px rgba(0,0,0,0.2);
        }
        .tile.blank {
            background: transparent;
            box-shadow: none;
        }
        .controls {
            margin: 20px;
        }
        button {
            background: linear-gradient(145deg, #4CAF50, #45a049);
            border: none;
            color: white;
            padding: 15px 32px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 4px 2px;
            cursor: pointer;
            border-radius: 8px;
            transition: all 0.3s ease;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 5px 15px rgba(0,0,0,0.3);
        }
        .step-info {
            background: rgba(255,255,255,0.1);
            padding: 15px;
            border-radius: 10px;
            margin: 20px 0;
            backdrop-filter: blur(10px);
        }
        .progress-bar {
            width: 100%;
            height: 20px;
            background: rgba(255,255,255,0.2);
            border-radius: 10px;
            overflow: hidden;
            margin: 20px 0;
        }
        .progress {
            height: 100%;
            background: linear-gradient(90deg, #4CAF50, #45a049);
            width: 0%;
            transition: width 0.5s ease;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🧩 8-Puzzle Solver Visualization</h1>
        
        <div class="step-info">
            <h2 id="stepTitle">Step 1: Initial State</h2>
            <p id="actionText">Starting configuration</p>
        </div>
        
        <div class="progress-bar">
            <div class="progress" id="progressBar"></div>
        </div>
        
        <div class="puzzle-board" id="puzzleBoard">
            <!-- Tiles will be generated by JavaScript -->
        </div>
        
        <div class="controls">
            <button onclick="prevStep()">⬅️ Previous</button>
            <button onclick="playPause()">▶️ Play</button>
            <button onclick="nextStep()">➡️ Next</button>
            <button onclick="reset()">🔄 Reset</button>
        </div>
        
        <div class="step-info">
            <p>Step <span id="currentStep">1</span> of <span id="totalSteps">` + fmt.Sprintf("%d", len(keyVals)) + `</span></p>
        </div>
    </div>

    <script>
        const steps = ` + pv.generateJSONSteps(keyVals, actions) + `;
        let currentStepIndex = 0;
        let isPlaying = false;
        let playInterval;

        function generateBoard(boardString) {
            const board = document.getElementById('puzzleBoard');
            board.innerHTML = '';
            
            for (let i = 0; i < 9; i++) {
                const tile = document.createElement('div');
                tile.className = 'tile';
                
                if (boardString[i] === 'B') {
                    tile.className += ' blank';
                    tile.textContent = '';
                } else {
                    tile.textContent = boardString[i];
                }
                
                board.appendChild(tile);
            }
        }

        function updateStep() {
            const step = steps[currentStepIndex];
            document.getElementById('stepTitle').textContent = 'Step ' + (currentStepIndex + 1) + ': ' + (step.action || 'Initial State');
            document.getElementById('actionText').textContent = step.action || 'Starting configuration';
            document.getElementById('currentStep').textContent = currentStepIndex + 1;
            
            const progress = ((currentStepIndex + 1) / steps.length) * 100;
            document.getElementById('progressBar').style.width = progress + '%';
            
            generateBoard(step.board);
        }

        function nextStep() {
            if (currentStepIndex < steps.length - 1) {
                currentStepIndex++;
                updateStep();
            }
        }

        function prevStep() {
            if (currentStepIndex > 0) {
                currentStepIndex--;
                updateStep();
            }
        }

        function playPause() {
            if (isPlaying) {
                clearInterval(playInterval);
                isPlaying = false;
                document.querySelector('button[onclick="playPause()"]').textContent = '▶️ Play';
            } else {
                playInterval = setInterval(() => {
                    if (currentStepIndex < steps.length - 1) {
                        nextStep();
                    } else {
                        playPause(); // Stop when reached the end
                    }
                }, 1000);
                isPlaying = true;
                document.querySelector('button[onclick="playPause()"]').textContent = '⏸️ Pause';
            }
        }

        function reset() {
            currentStepIndex = 0;
            updateStep();
            if (isPlaying) {
                playPause();
            }
        }

        // Initialize
        updateStep();
    </script>
</body>
</html>`

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating HTML file: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(htmlContent)
	fmt.Printf("Web visualization saved to: %s\n", filename)
	fmt.Println("Open this file in your web browser to see the animated solution!")
}

func (pv *PuzzleVisualizer) generateJSONSteps(keyVals []string, actions []string) string {
	result := "[\n"
	for i, key := range keyVals {
		action := ""
		if i > 0 && i-1 < len(actions) {
			action = actions[i-1]
		}

		result += fmt.Sprintf(`    {"board": "%s", "action": "%s"}`, key, action)
		if i < len(keyVals)-1 {
			result += ","
		}
		result += "\n"
	}
	result += "]"
	return result
}
