package main

import (
	"fmt"
	"time"
)

func main() {
	playerOneScore := 1
	playerTwoScore := 1
	playerOneScore := 0
	playerTwoScore := 0
  gameOver := false
  twoPlayer := false
  gameNumber := 0

  header := "           "
  playerOneWin := false
  playerTwoWin := false
	boardAvailableMoves := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	playerOneMoves := make([]int, 1, 10)
	playerTwoMoves := make([]int, 1, 10)

	var option string

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
	}
}

func generateLogic(playerOneMoves []int, playerTwoMoves []int, boardAvailableMoves []int) {

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
sq1 := "1"
sq2 := "2"
sq3 := "3"
sq4 := "4"
sq5 := "5"
sq6 := "6"
sq7 := "7"
sq8 := "8"
sq9 := "9"
