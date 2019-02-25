package ticktacktoe

import "fmt"

// Welcome melhor função do mundo
func Welcome() {
	fmt.Println("Welcome to the TicTacToe Game!")
	fmt.Println()
	game()
}

func getParameters() (int, int) {
	var choiceLine, choiceColumn int
	fmt.Scan(&choiceLine)
	fmt.Scan(&choiceColumn)

	switch {
	case choiceLine < 0 || choiceLine > 9:
		fmt.Println("Valor para linha incorreto!")
		return getParameters()
	case choiceColumn < 0 || choiceColumn > 9:
		fmt.Println("Valor para coluna incorreto!")
		return getParameters()
	}

	return choiceLine, choiceColumn
}

func printBoard(board [3][3]int) {
	fmt.Println("  _ _ _ ")
	for i := 0; i < 3; i++ {
		fmt.Print(i, "|")
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case 0:
				fmt.Print("_|")
			case 1:
				fmt.Print("X|")
			case 2:
				fmt.Print("O|")
			}
		}
		fmt.Println("")
	}
	fmt.Println("  0 1 2")
}

func game() {
	// Rodada
	var round = 0

	// Tabuleiro
	var board [3][3]int
	for {
		fmt.Println("Choice a line and a column by writing two numbers separated by space:")
		printBoard(board)

		var x, y = getParameters()
		board[x][y] = round%2 + 1
		round++

		if round == 9 {
			fmt.Println("Empate!")
			return
		}

		for i := 0; i < 3; i++ {
			if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
				winner(board[i][0])
				break
			} else if board[0][i] != 0 && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
				winner(board[0][i])
				break
			}
		}

		if board[1][1] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] ||
			board[1][1] != 0 && board[2][0] == board[1][1] && board[1][1] == board[0][2] {
			winner(board[1][1])
			break
		}
	}
}

func winner(v int) {
	var p string
	switch v {
	case 1:
		p = "X"
	case 2:
		p = "Y"
	}
	fmt.Println("Parabéns! Jogador ", p, " venceu!")
}
