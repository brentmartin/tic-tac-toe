package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int = 0
	var o int = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		println("    Tic-tac-toe    ")
		println("   Golang Edition  ")
		println("╔════════╦════════╗")
		println("║ X =", x, " ║ O =", o, " ║")
		println("╚════════╩════════╝")
		println("   ┍━━━━━━━━━━━┑   ")
		println("   ⎟ 1 ║ 2 ║ 3 ⎟   ")
		println("   ⎟═══╬═══╬═══⎟   ")
		println("   ⎟ 4 ║ 5 ║ 6 ⎟   ")
		println("   ⎟═══╬═══╬═══⎟   ")
		println("   ⎟ 7 ║ 8 ║ 9 ⎟   ")
		println("   ┕━━━━━━━━━━━┙   ")
		fmt.Println(scanner.Text())
	}
func generateLogic(playerOneMoves []int, playerTwoMoves []int, boardAvailableMoves []int) {

}

func generateBoard(playerOneScore, playerTwoScore) {
}
