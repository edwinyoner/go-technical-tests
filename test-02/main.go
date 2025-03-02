/*
Archivo: main.go
Autor: Edwin Yoner
Fecha: 11-11-2024
Versión: 1.0

Descripción:
Este programa solicita al usuario un número entero en el rango de [0 - 10] y calcula su factorial utilizando dos métodos:
uno iterativo y otro recursivo. El programa valida la entrada para asegurarse de que el usuario ingrese un número
entero en el rango permitido. Luego, muestra el resultado del cálculo factorial utilizando ambos métodos.
*/

package main

import (
	"bufio"   // Paquete para leer entradas del usuario de manera eficiente
	"fmt"     // Paquete para imprimir mensajes en la consola
	"os"      // Paquete para interactuar con el sistema operativo, en este caso para la entrada de datos
	"strconv" // Paquete para convertir strings a otros tipos de datos
	"strings" // Paquete para manipulación de strings, utilizado aquí para limpiar espacios en blanco
)

// Función principal del programa
func main() {
	scanner := bufio.NewScanner(os.Stdin) // Crea un scanner para leer la entrada del usuario
	var numero int                        // Variable para almacenar el número ingresado

	for {
		fmt.Println("Ingrese un número entero en el rango de [0 - 10]:") // Solicita al usuario un número
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text()) // Lee y limpia la entrada del usuario

		// Convierte el input a un entero y valida la entrada
		num, err := strconv.Atoi(input)
		if err != nil { // Si ocurre un error en la conversión
			fmt.Printf("Error: ingresa un número válido\n")
			continue
		}

		// Verifica que el número esté dentro del rango permitido
		if num < 0 || num > 10 {
			fmt.Println("Error: el número debe estar en el rango de [0 - 10]")
			continue
		}

		numero = num // Asigna el número válido a la variable
		// Muestra el resultado del factorial calculado por ambos métodos
		fmt.Printf("El factorial de %d (iterativo) es: %d\n", numero, calculateFactorial(numero))
		fmt.Printf("El factorial de %d (recursivo) es: %d\n", numero, calculateFactorialRecursivo(numero))
		break
	}
}

// calculateFactorial calcula el factorial de un número entero usando un método iterativo
// Parámetro:
// - numero: el número entero del cual calcular el factorial
// Retorno:
// - int: el resultado del cálculo factorial
func calculateFactorial(numero int) int {
	if numero == 0 { // El factorial de 0 es 1 por definición
		return 1
	}

	factorial := 1
	for i := 1; i <= numero; i++ { // Bucle para multiplicar cada número hasta 'numero'
		factorial *= i
	}

	return factorial // Retorna el resultado del factorial
}

// calculateFactorialRecursivo calcula el factorial de un número entero usando un método recursivo
// Parámetro:
// - numero: el número entero del cual calcular el factorial
// Retorno:
// - int: el resultado del cálculo factorial
func calculateFactorialRecursivo(numero int) int {
	if numero == 0 { // Condición base: el factorial de 0 es 1
		return 1
	}

	return numero * calculateFactorialRecursivo(numero-1) // Llamada recursiva para calcular el factorial
}
