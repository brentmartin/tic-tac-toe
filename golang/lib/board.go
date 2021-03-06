package lib

import "fmt"

var sq []string
var pOneScore int
var pTwoScore int
var header string

func init() {
	sq = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	pOneScore = 0
	pTwoScore = 0
	header = "   Golang Edition  "
}

func showBoard() {
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
}
