package main

import (
	"fmt"
	"bufio"
	"os"
	. "../model"
)

func main() {
	// entry
	scanner := bufio.NewScanner(os.Stdin)

	// my tictac
	var myTicTac TicTac
	myTicTac.Plateau = [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	myTicTac.Turn = 1

	fmt.Print("Enter player 1: ")
	scanner.Scan()
	myTicTac.Player1 = scanner.Text()

	fmt.Print("Enter player 2: ")
	scanner.Scan()
	myTicTac.Player2 = scanner.Text()

	fmt.Printf("%v vs %v \n\n", myTicTac.Player1, myTicTac.Player2)
	var selector = switchTurn(myTicTac)

	fmt.Print("Possible value for row && col \n")
	fmt.Print("1 2 3 \n")
	fmt.Print("2 _ _ \n")
	fmt.Print("3 _ _ \n\n")

	for {
		var row, col int

		row = askSelection("row")
		col = askSelection("col")

		if isValid(myTicTac.Plateau[row][col]) {
			myTicTac.Plateau[row][col] = selector
			show(myTicTac)

			if hasWin(myTicTac, selector) {
				fmt.Printf("\n BRAVO %s, you Win !\n", returnUser(myTicTac))
				os.Exit(2)
			}
			if complete(myTicTac) {
				fmt.Printf("\n Equality !\n")
				os.Exit(2)
			}

			myTicTac.Turn ++
			selector = switchTurn(myTicTac)
		} else {
			fmt.Print("\n INVALID SELECTION \n")
		}
	}
}

func returnUser(t TicTac) string {
	if t.Turn%2 != 0 {
		return t.Player1
	} else {
		return t.Player2
	}
}

func askSelection(s string) int {
	var goodInt int
	fmt.Printf("selection : %s =", s)
	fmt.Scanf("%d ", &goodInt)
	if goodInt == 1 || goodInt == 2 || goodInt == 3 {
		return goodInt - 1
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
	if t.Turn%2 != 0 {
		return "o"
	} else {
		return "x"
	}
}

func complete(t TicTac) bool {
	for i := 0; i < len(t.Plateau); i++ {
		for j := 0; j < len(t.Plateau[i]); j++ {
			if t.Plateau[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func show(t TicTac) bool {
	for i := 0; i < len(t.Plateau); i++ {
		fmt.Print("[ ")
		for j := 0; j < len(t.Plateau[i]); j++ {
			fmt.Print(t.Plateau[i][j])
			fmt.Print(" ")
		}
		fmt.Print("]")
		fmt.Print("\n")
	}
	return true
}

func hasWin(t TicTac, char string) bool {
	plateau := t.Plateau
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
