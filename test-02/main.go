package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int

	for {
		fmt.Println("Ingrese un número entero en rango de [0 - 10]:")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: ingresa un número válido")
			continue
		}
		if num < 0 || num > 10 {
			fmt.Println("Error: el número debe estar en el rango de [0 - 10]")
			continue
		}
		numero = num
		fmt.Printf("El factorial de %d es: %d\n", numero, calculateFactorial(numero))
		break
	}
}

func calculateFactorial(numero int) int {
	if numero == 0 {
		return 1
	}

	factorial := 1
	for i := 1; i <= numero; i++ {
		factorial *= i
	}

	return factorial
}
