package main

import (
	"fmt"
	"time"
)

func main() {
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

	for i := 1; i <= 3; i++ {
		generateBoard(gameOn, sq, pOneMoves)

		if gameOn == false {
			fmt.Println("C) Play against Computer")
			fmt.Println("P) Play against Player")
		}

		if gameOn == true {
			fmt.Println("The numbers correspond with the tiles")
			fmt.Println("Choose a number 1-9 to make your move")
		}

		fmt.Print("Please choose an option: ")
		fmt.Scanln(&choice)

		switch {
		case choice == "1" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 1)
		case choice == "2" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 2)
		case choice == "3" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 3)
		case choice == "4" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 4)
		case choice == "5" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 5)
		case choice == "6" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 6)
		case choice == "7" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 7)
		case choice == "8" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 8)
		case choice == "9" && gameOn == true:
			doThings(pOneMoves, boardMoves, sq, 9)
		default:
			fmt.Println("Try again plz")
		}
	}
}

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
