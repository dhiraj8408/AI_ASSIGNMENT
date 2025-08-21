package main

import (
	"fmt"
	"os"
	"strings"
)

// GenerateWebVisualization creates an HTML file for web-based puzzle visualization
func GenerateWebVisualization(keyVals []string, actions []string, algorithmName string, nodesExpanded int, uiEnabled bool) {
	if !uiEnabled {
		return
	}

	filename := fmt.Sprintf("puzzle_solution_%s.html", strings.ToLower(algorithmName))

	htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>8-Puzzle Solver - %s Visualization</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            margin: 0;
            padding: 20px;
            color: white;
            min-height: 100vh;
        }
        .container {
            max-width: 1000px;
            margin: 0 auto;
            text-align: center;
        }
        .header {
            background: rgba(255,255,255,0.1);
            padding: 20px;
            border-radius: 15px;
            margin-bottom: 20px;
            backdrop-filter: blur(10px);
        }
        .algorithm-name {
            font-size: 2.5em;
            font-weight: bold;
            color: #FFD700;
            text-shadow: 2px 2px 4px rgba(0,0,0,0.5);
            margin-bottom: 10px;
        }
        .stats {
            display: flex;
            justify-content: center;
            gap: 30px;
            margin: 20px 0;
        }
        .stat-item {
            background: rgba(255,255,255,0.2);
            padding: 15px 25px;
            border-radius: 10px;
            backdrop-filter: blur(5px);
        }
        .stat-number {
            font-size: 1.8em;
            font-weight: bold;
            color: #FFD700;
        }
        .stat-label {
            font-size: 0.9em;
            opacity: 0.8;
        }
        .puzzle-board {
            display: inline-grid;
            grid-template-columns: repeat(3, 100px);
            grid-gap: 8px;
            background: rgba(0,0,0,0.3);
            padding: 25px;
            border-radius: 20px;
            margin: 20px;
            box-shadow: 0 15px 35px rgba(0,0,0,0.4);
            backdrop-filter: blur(10px);
        }
        .tile {
            width: 100px;
            height: 100px;
            background: linear-gradient(145deg, #f0f0f0, #d0d0d0);
            border-radius: 15px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 32px;
            font-weight: bold;
            color: #333;
            transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            box-shadow: 0 8px 16px rgba(0,0,0,0.3);
            position: relative;
            overflow: hidden;
        }
        .tile.blank {
            background: transparent;
            box-shadow: inset 0 0 20px rgba(255,255,255,0.1);
            border: 2px dashed rgba(255,255,255,0.3);
        }
        .tile.highlight {
            transform: scale(1.1);
            background: linear-gradient(145deg, #FFD700, #FFA500);
            color: #333;
            box-shadow: 0 0 25px rgba(255, 215, 0, 0.6);
        }
        .tile::before {
            content: '';
            position: absolute;
            top: -50%%;
            left: -50%%;
            width: 200%%;
            height: 200%%;
            background: linear-gradient(45deg, transparent, rgba(255,255,255,0.3), transparent);
            transform: rotate(45deg);
            transition: all 0.6s;
            opacity: 0;
        }
        .tile:hover::before {
            opacity: 1;
            animation: shine 0.6s ease-in-out;
        }
        @keyframes shine {
            0%% { transform: translateX(-100%%) translateY(-100%%) rotate(45deg); }
            100%% { transform: translateX(100%%) translateY(100%%) rotate(45deg); }
        }
        .controls {
            margin: 30px 0;
            display: flex;
            justify-content: center;
            gap: 15px;
            flex-wrap: wrap;
        }
        button {
            background: linear-gradient(145deg, #4CAF50, #45a049);
            border: none;
            color: white;
            padding: 15px 30px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            font-weight: bold;
            cursor: pointer;
            border-radius: 50px;
            transition: all 0.3s ease;
            box-shadow: 0 4px 15px rgba(0,0,0,0.2);
            min-width: 120px;
        }
        button:hover {
            transform: translateY(-3px);
            box-shadow: 0 8px 25px rgba(0,0,0,0.3);
            background: linear-gradient(145deg, #5CBF60, #4CAF50);
        }
        button:active {
            transform: translateY(-1px);
        }
        button.play-pause {
            background: linear-gradient(145deg, #FF6B6B, #FF5252);
        }
        button.play-pause:hover {
            background: linear-gradient(145deg, #FF7B7B, #FF6B6B);
        }
        button.reset {
            background: linear-gradient(145deg, #9C27B0, #8E24AA);
        }
        button.reset:hover {
            background: linear-gradient(145deg, #AB47BC, #9C27B0);
        }
        .step-info {
            background: rgba(255,255,255,0.15);
            padding: 20px;
            border-radius: 15px;
            margin: 20px 0;
            backdrop-filter: blur(15px);
            border: 1px solid rgba(255,255,255,0.2);
        }
        .step-title {
            font-size: 1.5em;
            font-weight: bold;
            color: #FFD700;
            margin-bottom: 10px;
        }
        .action-text {
            font-size: 1.1em;
            background: rgba(0,0,0,0.2);
            padding: 10px 15px;
            border-radius: 8px;
            display: inline-block;
            margin: 5px 0;
        }
        .progress-container {
            background: rgba(255,255,255,0.1);
            border-radius: 25px;
            padding: 8px;
            margin: 20px 0;
        }
        .progress-bar {
            width: 100%%;
            height: 30px;
            background: rgba(255,255,255,0.2);
            border-radius: 20px;
            overflow: hidden;
            position: relative;
        }
        .progress {
            height: 100%%;
            background: linear-gradient(90deg, #4CAF50, #45a049, #66BB6A);
            width: 0%%;
            transition: width 0.5s ease;
            border-radius: 20px;
            position: relative;
        }
        .progress::after {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            height: 100%%;
            width: 100%%;
            background: linear-gradient(90deg, transparent, rgba(255,255,255,0.4), transparent);
            animation: progressShine 2s infinite;
        }
        @keyframes progressShine {
            0%% { transform: translateX(-100%%); }
            100%% { transform: translateX(100%%); }
        }
        .progress-text {
            text-align: center;
            margin-top: 10px;
            font-weight: bold;
        }
        .footer {
            margin-top: 40px;
            padding: 20px;
            background: rgba(0,0,0,0.2);
            border-radius: 10px;
            font-size: 0.9em;
            opacity: 0.8;
        }
        @media (max-width: 768px) {
            .puzzle-board {
                grid-template-columns: repeat(3, 80px);
            }
            .tile {
                width: 80px;
                height: 80px;
                font-size: 24px;
            }
            .controls {
                flex-direction: column;
                align-items: center;
            }
            .stats {
                flex-direction: column;
                gap: 15px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <div class="algorithm-name">%s Algorithm</div>
            <div class="stats">
                <div class="stat-item">
                    <div class="stat-number">%d</div>
                    <div class="stat-label">Steps in Solution</div>
                </div>
                <div class="stat-item">
                    <div class="stat-number">%d</div>
                    <div class="stat-label">Nodes Expanded</div>
                </div>
                <div class="stat-item">
                    <div class="stat-number">%.2f</div>
                    <div class="stat-label">Efficiency Ratio</div>
                </div>
            </div>
        </div>
        
        <div class="step-info">
            <div class="step-title" id="stepTitle">Step 1: Initial State</div>
            <div class="action-text" id="actionText">Starting configuration</div>
        </div>
        
        <div class="progress-container">
            <div class="progress-bar">
                <div class="progress" id="progressBar"></div>
            </div>
            <div class="progress-text">
                Step <span id="currentStep">1</span> of <span id="totalSteps">%d</span>
            </div>
        </div>
        
        <div class="puzzle-board" id="puzzleBoard">
            <!-- Tiles will be generated by JavaScript -->
        </div>
        
        <div class="controls">
            <button onclick="firstStep()">⏮️ First</button>
            <button onclick="prevStep()">⬅️ Previous</button>
            <button onclick="playPause()" class="play-pause">▶️ Play</button>
            <button onclick="nextStep()">➡️ Next</button>
            <button onclick="lastStep()">⏭️ Last</button>
            <button onclick="reset()" class="reset">🔄 Reset</button>
        </div>
        
        <div class="footer">
            <p><strong>%s Algorithm Visualization</strong></p>
            <p>Generated automatically by 8-Puzzle Solver</p>
            <p>Use the controls above to navigate through the solution steps</p>
        </div>
    </div>

    <script>
        const steps = %s;
        let currentStepIndex = 0;
        let isPlaying = false;
        let playInterval;
        let playSpeed = 1000; // milliseconds

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
                
                // Add highlight effect for recently moved tiles
                if (currentStepIndex > 0) {
                    const prevBoard = steps[currentStepIndex - 1].board;
                    if (prevBoard[i] !== boardString[i] && boardString[i] !== 'B') {
                        tile.className += ' highlight';
                    }
                }
                
                board.appendChild(tile);
            }
        }

        function updateStep() {
            const step = steps[currentStepIndex];
            const stepTitle = document.getElementById('stepTitle');
            const actionText = document.getElementById('actionText');
            
            stepTitle.textContent = 'Step ' + (currentStepIndex + 1) + ': ' + (currentStepIndex === 0 ? 'Initial State' : 'Move ' + currentStepIndex);
            actionText.textContent = step.action || 'Starting configuration';
            
            document.getElementById('currentStep').textContent = currentStepIndex + 1;
            
            const progress = ((currentStepIndex + 1) / steps.length) * 100;
            document.getElementById('progressBar').style.width = progress + '%%';
            
            generateBoard(step.board);
        }

        function nextStep() {
            if (currentStepIndex < steps.length - 1) {
                currentStepIndex++;
                updateStep();
            } else if (isPlaying) {
                playPause(); // Stop auto-play when reached the end
            }
        }

        function prevStep() {
            if (currentStepIndex > 0) {
                currentStepIndex--;
                updateStep();
            }
        }

        function firstStep() {
            currentStepIndex = 0;
            updateStep();
            if (isPlaying) playPause();
        }

        function lastStep() {
            currentStepIndex = steps.length - 1;
            updateStep();
            if (isPlaying) playPause();
        }

        function playPause() {
            const button = document.querySelector('.play-pause');
            if (isPlaying) {
                clearInterval(playInterval);
                isPlaying = false;
                button.textContent = '▶️ Play';
            } else {
                playInterval = setInterval(() => {
                    if (currentStepIndex < steps.length - 1) {
                        nextStep();
                    } else {
                        playPause(); // Stop when reached the end
                    }
                }, playSpeed);
                isPlaying = true;
                button.textContent = '⏸️ Pause';
            }
        }

        function reset() {
            currentStepIndex = 0;
            updateStep();
            if (isPlaying) {
                playPause();
            }
        }

        // Keyboard shortcuts
        document.addEventListener('keydown', function(event) {
            switch(event.key) {
                case 'ArrowLeft':
                    prevStep();
                    break;
                case 'ArrowRight':
                    nextStep();
                    break;
                case ' ':
                    event.preventDefault();
                    playPause();
                    break;
                case 'Home':
                    firstStep();
                    break;
                case 'End':
                    lastStep();
                    break;
                case 'r':
                case 'R':
                    reset();
                    break;
            }
        });

        // Initialize
        updateStep();
        
        // Auto-start animation after 2 seconds
        setTimeout(() => {
            if (!isPlaying && steps.length > 1) {
                playPause();
            }
        }, 2000);
    </script>
</body>
</html>`, algorithmName, algorithmName, len(keyVals), nodesExpanded, calculateEfficiency(len(keyVals), nodesExpanded), len(keyVals), algorithmName, generateJSONSteps(keyVals, actions))

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating HTML file: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(htmlContent)
	fmt.Printf("🌐 Web visualization saved to: %s\n", filename)
	fmt.Printf("   Open this file in your web browser to see the animated solution!\n")
}

func calculateEfficiency(steps, nodesExpanded int) float64 {
	if nodesExpanded == 0 {
		return 0.0
	}
	return float64(steps) / float64(nodesExpanded) * 100
}

func generateJSONSteps(keyVals []string, actions []string) string {
	if len(keyVals) == 0 {
		return "[]"
	}

	result := "[\n"
	for i, key := range keyVals {
		action := ""
		if i > 0 && i-1 < len(actions) {
			action = actions[i-1]
		}

		result += fmt.Sprintf(`        {"board": "%s", "action": "%s"}`, key, action)
		if i < len(keyVals)-1 {
			result += ","
		}
		result += "\n"
	}
	result += "    ]"
	return result
}
