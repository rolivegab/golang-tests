package calculator

import "fmt"

func askChoice() {
	for {
		var choice int
		fmt.Println("Digite uma opção:")
		fmt.Println("1- Somar")
		fmt.Println("2- Subtrair")
		fmt.Println("3- Multiplicação")
		fmt.Println("4- Dividir")

		fmt.Scan(&choice)
		fmt.Println("Você escolheu", choice)

		switch choice {
		case 1:
			soma()
			break
		case 2:
			subtrair()
			break
		case 3:
			multiplicar()
			break
		case 4:
			dividir()
			break
		default:
			fmt.Println("Opção incorreta! Por favor, tente novamente.")
			askChoice()
		}
	}
}

func getParameters() (int, int) {
	var a, b int
	fmt.Println("Digite os dois números separados por espaço, e em seguida, aperte ENTER.")
	fmt.Scan(&a)
	fmt.Scan(&b)
	return a, b
}

func soma() {
	var a, b = getParameters()
	fmt.Println("Resultado:", a+b)
}

func subtrair() {
	var a, b = getParameters()
	fmt.Println("Resultado:", a-b)
}

func multiplicar() {
	var a, b = getParameters()
	fmt.Println("Resultado:", a*b)
}

func dividir() {
	var a, b = getParameters()
	fmt.Println("Resultado:", a/b)
}

// Run melhor função desse arquivo, hehehe
func Run() {
	welcome()
	askChoice()
}

func welcome() {
	fmt.Println("Bem-vindo à calculadora")
}
