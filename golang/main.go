package main

import (
	"fmt"
	// "time"
)

var pOneScore int
var pTwoScore int
var header string
var choice string
// var gameNumber int
var boardMoves []int
var pOneMoves []int

// var pTwoMoves []int
var gameOn bool
var sq []string

// var pOneWin bool
// var pTwoWin bool
// var twoPlayer bool
func main() {
	header = "   Golang Edition  "
	sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	boardMoves = make([]int, 1, 10)
	pOneMoves = make([]int, 1, 10)
	// pTwoMoves = make([]int, 1, 10)
	pOneScore = 0
	pTwoScore = 0
	gameOn = true

func boardUpdater() {


	for i := 1; i <= 10; i++ {

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
			doThings(1)
		case choice == "2" && gameOn == true:
			doThings(2)
		case choice == "3" && gameOn == true:
			doThings(3)
		case choice == "4" && gameOn == true:
			doThings(4)
		case choice == "5" && gameOn == true:
			doThings(5)
		case choice == "6" && gameOn == true:
			doThings(6)
		case choice == "7" && gameOn == true:
			doThings(7)
		case choice == "8" && gameOn == true:
			doThings(8)
		case choice == "9" && gameOn == true:
			doThings(9)
		default:
			fmt.Println("Try again plz")
		}
	}
}

	if gameOn == false {
		pOneMoves = make([]int, 1, 10)
		sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		gameOn = true
	}

	fmt.Println("")
	fmt.Println("    Tic-tac-toe    ")
	fmt.Println(header)
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

func doThings(x int) {
	pOneMoves = append(pOneMoves, x)
	boardMoves = append(boardMoves, x)

	for i := 1; i <= 9; i++ {
		if integerInSlice(i, pOneMoves) {
			sq[i] = "X"
		}
	}
}

// time.Sleep(time.Millisecond * 2000)

