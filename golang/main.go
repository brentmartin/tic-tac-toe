package main

import (
	"fmt"
	"time"
)

func main() {


	fmt.Println("1) Play against Computer")
	fmt.Println("2) Play against Player")
	fmt.Print("Please choose an option: ")

	fmt.Scanln(&option)

	switch option {
	case "1":
		generateLogic(playerOneMoves, playerTwoMoves, boardAvailableMoves)
		generateBoard(playerOneScore, playerTwoScore)
	case "2":
		generateBoard(playerOneScore, playerTwoScore)
	default:
		fmt.Println("Uknown option, no action taken")
	pOneScore := 0
	pTwoScore := 0
	gameOn := false
	twoPlayer := false
	gameNumber := 0

	header := "           "

	sq := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	pOneWin := false
	pTwoWin := false
	boardMoves := make([]int, 1, 10)
	pOneMoves := make([]int, 1, 10)
	pTwoMoves := make([]int, 1, 10)

	var choice string
	}
}

func generateLogic(playerOneMoves []int, playerTwoMoves []int, boardAvailableMoves []int) {
func generateBoard(gameOn bool, sq []string, pOneMoves []int) {

}

func generateBoard(playerOneScore int, playerTwoScore int) {
	playerOneScore = +1
	playerTwoScore = +1
	println(playerOneScore)
	println(playerTwoScore)
	println("")
	println("    Tic-tac-toe    ")
	println("   Golang Edition  ")
	println("╔════════╦════════╗")
	println("║ X =", playerOneScore, " ║ O =", playerTwoScore, " ║")
	println("╚════════╩════════╝")
	println("   ┍━━━━━━━━━━━┑   ")
	println("   ⎟ ", sq1," ║ ", sq2," ║ ", sq3," ⎟   ")
	println("   ⎟═══╬═══╬═══⎟   ")
	println("   ⎟ ", sq4," ║ ", sq5," ║ ", sq6," ⎟   ")
	println("   ⎟═══╬═══╬═══⎟   ")
	println("   ⎟ ", sq7," ║ ", sq8," ║ ", sq9," ⎟   ")
	println("   ┕━━━━━━━━━━━┙   ")
}
