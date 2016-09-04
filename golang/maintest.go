package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

var move = make(chan int)

type Player struct {
	score int
	moves []int
	// win   bool
	// turn  bool
}

var p1 Player

// var p1 = Player{
//   score: 0,
//   moves: make([]int, 1, 10),
// }

var pOneScore int
var pTwoScore int

// var pOneWin bool
// var pTwoWin bool
// var twoPlayer bool
var pOneMoves []int

// var pTwoMoves []int
var boardMoves []int

var header string
var choice string

// var gameNumber int

var gameOn bool
var sq []string

var m int

var choiceString string
var choiceInt int

var found bool
var playUnavail bool

func main() {
	header = "   Golang Edition  "
	sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	boardMoves = make([]int, 1, 10)
	pOneMoves = make([]int, 1, 10)
	// pTwoMoves = make([]int, 1, 10)
	pOneScore = 0
	pTwoScore = 0
	gameOn = true

	p1 := Player{
		score: 0,
		moves: make([]int, 1, 10),
	}

	playUnavail = false

	wg.Add(2)
	go playController()
	go boardUpdater()
	wg.Wait()

}

func boardUpdater() {

	for {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("")
		fmt.Println("    Tic-tac-toe    ")
		fmt.Println(header)
		fmt.Println("╔════════╦════════╗")
		fmt.Println("║ X =", pOneScore, " ║ O =", pTwoScore, " ║")
		fmt.Println("╚════════╩════════╝")
		fmt.Println("   ┍━━━━━━━━━━━┑   ")
		fmt.Println("   ⎟", sq[1], "║", sq[2], "║", sq[3], "⎟   ")
		fmt.Println("   ⎟═══╬═══╬═══⎟   ")
		fmt.Println("   ⎟", sq[4], "║", sq[5], "║", sq[6], "⎟   ")
		fmt.Println("   ⎟═══╬═══╬═══⎟   ")
		fmt.Println("   ⎟", sq[7], "║", sq[8], "║", sq[9], "⎟   ")
		fmt.Println("   ┕━━━━━━━━━━━┙   ")
		fmt.Println("")

		if m == 101 {
			fmt.Println("        Tile already selected!       ")
			fmt.Println("Please choose an eligible number play")
		}
		if m == 102 {
			fmt.Println("                                     ")
			fmt.Println("Please choose an eligible number play")
		}
		m = <-move
		if m >= 1 && m <= 9 {
			doThings(m)
		}
	}
}

func playController() {
	for {
		if gameOn == false {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("C) Play against Computer")
			fmt.Println("P) Play against Player")
		}

		if gameOn == true && playUnavail == false {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("The numbers correspond with the tiles")
			fmt.Println("Choose a number 1-9 to make your move")
		}

		if gameOn == true && playUnavail == true {
			move <- 101
			time.Sleep(time.Millisecond * 500)
			move <- 102
			time.Sleep(time.Millisecond * 500)
			move <- 101
			time.Sleep(time.Millisecond * 500)
			move <- 102
			time.Sleep(time.Millisecond * 500)
			move <- 101
			time.Sleep(time.Millisecond * 500)
			move <- 102
			time.Sleep(time.Millisecond * 500)
			move <- 0
			time.Sleep(time.Millisecond * 50)
			fmt.Println("The numbers correspond with the tiles")
			fmt.Println("Choose a number 1-9 to make your move")
		}

		fmt.Print("Please choose an option: ")
		fmt.Scanln(&choice)

		convStringToInt(choice)
		playUnavail = integerInSlice(choiceInt, boardMoves)

		// switch {
		// case playUnavail == true:
		// 	break
		// case choiceInt >= 1 && choiceInt <= 9:
		// 	move <- choiceInt
		// default:
		// 	fmt.Println("Try again plz")
		// }

		switch {
		case playUnavail == true:
			break
		case choice == "1" && gameOn == true:
			move <- 1
		case choice == "2" && gameOn == true:
			move <- 2
		case choice == "3" && gameOn == true:
			move <- 3
		case choice == "4" && gameOn == true:
			move <- 4
		case choice == "5" && gameOn == true:
			move <- 5
		case choice == "6" && gameOn == true:
			move <- 6
		case choice == "7" && gameOn == true:
			move <- 7
		case choice == "8" && gameOn == true:
			move <- 8
		case choice == "9" && gameOn == true:
			move <- 9
		default:
			fmt.Println("Try again plz")
		}
	}
}

func integerInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// // WITHOUT STRUCTS
// func doThings(m int) {
// 	pOneMoves = append(pOneMoves, m)
// 	boardMoves = append(boardMoves, m)
//
// 	for i := 1; i <= 9; i++ {
// 		if integerInSlice(i, boardMoves) {
// 			sq[i] = "X"
// 		}
// 	}
// }

// STRUCTS ON BOI!!!!
func doThings(m int) {
	p1.moves = append(playerOne.moves, m)
	boardMoves = append(boardMoves, m)

	for i := 1; i <= 9; i++ {
		if integerInSlice(i, boardMoves) {
			sq[i] = "X"
		}
	}
}

func convStringToInt(text string) {
	choiceString = text
	if s, err := strconv.Atoi(choiceString); err == nil {
		choiceInt = s
	}
}
