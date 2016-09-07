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
var check = make(chan int)

var pOneScore int
var pTwoScore int
var header string
var choice string
// var gameNumber int
var boardMoves []int

var pOneTurn bool
var pTwoTurn bool
var pOneWin bool
var pTwoWin bool
var pOneMoves []int
var move = make(chan int)

// var pTwoMoves []int
var gameOn bool
var sq []string

// var pOneWin bool
// var pTwoWin bool
// var twoPlayer bool
var m int
var c int

var choiceString string
var choiceInt int

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

	playUnavail = false

	wg.Add(3)
	go boardController()
	go boardInterface()
	go boardLogic()
	wg.Wait()

}
func boardUpdater() {

	// if gameOn == false {
	// 	pOneMoves = make([]int, 1, 10)
	// 	sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// 	gameOn = true
	// }

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
func boardLogic() {
	var turn int
	var slice []int
	var win bool
	for {
		c <- check

		if pOneTurn == true {
			slice = pOneMoves
			turn = 1
		}
		if pTwoTurn == true {
			slice = pTwoMoves
			turn = 2
		}

	}
}
	for {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
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
			check <- m
		}
	}
}

func boardController() {
	for {
		if gameOn == false {
			time.Sleep(time.Millisecond * 50)
			fmt.Println("C) Play against Computer")
			fmt.Println("P) Play against Player")
		}

		if gameOn == true {
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

func doThings(m int) {
	pOneMoves = append(pOneMoves, m)
	boardMoves = append(boardMoves, m)

	for i := 1; i <= 9; i++ {
		if integerInSlice(i, pOneMoves) {
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

func checkWin(a int, b int, c int, slice []int) bool {
	if integerInSlice(a, slice) {
		return true
	}
	return false
}
