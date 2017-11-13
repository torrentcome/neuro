package main

import (
	"fmt"
	"bufio"
	"os"
)

type TicTac struct {
	plateau [][] string
	turn    int
	player1 string
	player2 string
}

func main() {
	// entry
	scanner := bufio.NewScanner(os.Stdin)

	// my tictac
	var myTicTac TicTac
	myTicTac.plateau = [][]string{[]string{"_", "_", "_"}, []string{"_", "_", "_"}, []string{"_", "_", "_"}}
	myTicTac.turn = 1

	fmt.Print("Enter player 1: ")
	scanner.Scan()
	myTicTac.player1 = scanner.Text()

	fmt.Print("Enter player 2: ")
	scanner.Scan()
	myTicTac.player2 = scanner.Text()

	fmt.Printf("%v vs %v \n\n", myTicTac.player1, myTicTac.player2)
	var selector = switchTurn(myTicTac)

	fmt.Print("Possible value for row && col -> 0 1 2 \n\n")

	for {
		var row, col int

		row = askSelection("row")
		col = askSelection("col")

		if isValid(myTicTac.plateau[row][col]) {
			myTicTac.plateau[row][col] = selector
			show(myTicTac)

			if hasWin(myTicTac, selector) {
				fmt.Printf("\n BRAVO %s, you Win !\n", returnUser(myTicTac))
				os.Exit(2)
			}
			if complete(myTicTac) {
				fmt.Printf("\n Equality !\n")
				os.Exit(2)
			}

			myTicTac.turn ++
			selector = switchTurn(myTicTac)
		} else {
			fmt.Print("\n INVALID SELECTION \n")
		}
	}
}

func returnUser(t TicTac) string {
	if t.turn%2 != 0 {
		return t.player1
	} else {
		return t.player2
	}
}

func askSelection(s string) int {
	var goodInt int;
	fmt.Printf("selection : %s =", s)
	fmt.Scanf("%d ", &goodInt)
	if goodInt == 0 || goodInt == 1 || goodInt == 2 {
		return goodInt
	} else {
		return askSelection(s)
	}
}

func isValid(s string) bool {
	if s == "_" {
		return true
	} else
	{
		return false
	}
}

func switchTurn(t TicTac) string {
	if t.turn%2 != 0 {
		return "o"
	} else {
		return "x"
	}
}

func complete(t TicTac) bool {
	for i := 0; i < len(t.plateau); i++ {
		for j := 0; j < len(t.plateau[i]); j++ {
			if t.plateau[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func show(t TicTac) bool {
	for i := 0; i < len(t.plateau); i++ {
		fmt.Print("[ ")
		for j := 0; j < len(t.plateau[i]); j++ {
			fmt.Print(t.plateau[i][j])
			fmt.Print(" ")
		}
		fmt.Print("]")
		fmt.Print("\n")
	}
	return true
}

func hasWin(t TicTac, char string) bool {
	plateau := t.plateau
	// horizontal
	if (plateau[0][0] == char && plateau[0][1] == char && plateau[0][2] == char) ||
		(plateau[1][0] == char && plateau[1][1] == char && plateau[1][2] == char) ||
		(plateau[2][0] == char && plateau[2][1] == char && plateau[2][2] == char) ||
	// 	vertical
		(plateau[0][0] == char && plateau[1][0] == char && plateau[2][0] == char) ||
		(plateau[0][1] == char && plateau[1][1] == char && plateau[2][1] == char) ||
		(plateau[0][2] == char && plateau[1][2] == char && plateau[2][2] == char) ||
	// diagonal
		(plateau[0][0] == char && plateau[1][1] == char && plateau[2][2] == char) ||
		(plateau[0][2] == char && plateau[1][1] == char && plateau[2][0] == char) {
		return true
	} else
	{
		return false
	}
}
