package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// Read the three lines of the board
	var board [3]string
	for i := 0; i < 3; i++ {
		scanner.Scan()
		board[i] = scanner.Text()
	}

	// Validate the board
	if !isValidBoard(board) {
		fmt.Println("ERROR")
		return
	}

	// Count occurrences of 'X' and 'O'
	xCount, oCount := countSymbols(board)

	// Check for winners
	xWins := checkWinner(board, 'X')
	oWins := checkWinner(board, 'O')

	// Determine the game state
	if xWins && oWins {
		fmt.Println("ERROR")
	} else if xWins {
		fmt.Println("X WINS")
	} else if oWins {
		fmt.Println("O WINS")
	} else if xCount+oCount == 9 {
		fmt.Println("DRAW")
	} else if xCount > oCount {
		fmt.Println("O TURN")
	} else {
		fmt.Println("X TURN")
	}
}

// Validate the board
func isValidBoard(board [3]string) bool {
	for _, row := range board {
		if len(row) != 3 || !strings.ContainsAny(row, "XO.") || len(row) != len(strings.Trim(row, "XO.")) {
			return false
		}
	}
	return true
}

// Count occurrences of 'X' and 'O'
func countSymbols(board [3]string) (int, int) {
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, char := range row {
			if char == 'X' {
				xCount++
			} else if char == 'O' {
				oCount++
			}
		}
	}
	return xCount, oCount
}

// Check if a player has won
func checkWinner(board [3]string, player rune) bool {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if (rune(board[i][0]) == player && rune(board[i][1]) == player && rune(board[i][2]) == player) ||
			(rune(board[0][i]) == player && rune(board[1][i]) == player && rune(board[2][i]) == player) {
			return true
		}
	}
	// Check diagonals
	if (rune(board[0][0]) == player && rune(board[1][1]) == player && rune(board[2][2]) == player) ||
		(rune(board[0][2]) == player && rune(board[1][1]) == player && rune(board[2][0]) == player) {
		return true
	}
	return false
}
