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
		generateBoard(gameOn, sq, pOneMoves, header)

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

func generateBoard(gameOn bool, sq []string, pOneMoves []int, header string) {
	if gameOn == false {
		pOneMoves = make([]int, 1, 10)
		sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		gameOn = true
	}

	fmt.Println("")
	fmt.Println("    Tic-tac-toe    ")
	fmt.Println("   Golang Edition  ")
	fmt.Println("╔════════╦════════╗")
	fmt.Println("║ X = 0  ║ O = 0  ║")
	fmt.Println("╚════════╩════════╝")
	fmt.Println("   ┍━━━━━━━━━━━┑   ")
	fmt.Println("   ⎟", sq[1], "║", sq[2], "║", sq[3], "⎟   ")
	fmt.Println("   ⎟═══╬═══╬═══⎟   ")
	fmt.Println("   ⎟", sq[4], "║", sq[5], "║", sq[6], "⎟   ")
	fmt.Println("   ⎟═══╬═══╬═══⎟   ")
	fmt.Println("   ⎟", sq[7], "║", sq[8], "║", sq[9], "⎟   ")
	fmt.Println("   ┕━━━━━━━━━━━┙   ")
}

func integerInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func doThings(pM []int, bM []int, sq []string, x int) ([]int, []int, []string) {
	pM = append(pM, x)
	bM = append(bM, x)

	for i := 1; i <= 9; i++ {
		if integerInSlice(i, pM) {
			sq[i] = "X"
		}
	}

	return pM, bM, sq
}
