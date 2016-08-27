package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
  playerOneScore := int
	playerTwoScore := int
  boardAvailableMoves := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	playerOneMoves := []int
  playerTwoMoves := []int

  var option string

  fmt.Println("1) Play against Computer")
  fmt.Println("2) Play against Player")
  fmt.Print("Please choose an option: ")

  fmt.Scanln(&option)
func generateLogic(playerOneMoves []int, playerTwoMoves []int, boardAvailableMoves []int) {

}

func generateBoard(playerOneScore, playerTwoScore) {
}
