/*
Archivo: main.go
Autor: Edwin Yoner
Fecha: 10-11-2024
Versión: 1.0
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Función principal del programa. Pide al usuario ingresar una cadena,
// luego invoca las funciones reverseString y reverseStringPro para invertir la cadena
// y muestra los resultados en pantalla.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese una cadena por favor")

	// Escanear entrada del usuario y manejar posibles errores
	if !scanner.Scan() {
		fmt.Println("Error al leer la entrada:", scanner.Err())
		return
	}

	cadena := scanner.Text()

	// Validación de entrada para asegurar que la cadena no esté vacía
	if len(cadena) == 0 {
		fmt.Println("La cadena ingresada está vacía. Por favor, ingrese una cadena válida.")
		return
	}

	// Imprimir resultado de reverseString
	fmt.Println("Invertido con reverseString:", reverseString(cadena))

	// Imprimir resultado de reverseStringPro
	fmt.Println("Invertido con reverseStringPro:", reverseStringPro(cadena))
}

// reverseString invierte una cadena utilizando un bucle y construyendo
// una nueva cadena a partir de los caracteres originales.
// Parámetro:
//   - cadena: una cadena de texto para invertir.
//
// Retorna:
//   - La cadena invertida.
func reverseString(cadena string) string {
	r := []rune(cadena)               // Convierte la cadena a una lista de runas para soportar UTF-8.
	longitud := len(r) - 1            // Calcula la longitud para el índice inverso.
	resultado := make([]rune, len(r)) // Inicializa un slice de runas para almacenar el resultado.

	// Bucle para construir la cadena invertida.
	for i := 0; i < len(r); i++ {
		resultado[i] = r[longitud]
		longitud--
	}

	return string(resultado)
}

// reverseStringPro invierte una cadena de manera eficiente mediante un intercambio de runas.
// Este enfoque es más directo y realiza la inversión in-place.
// Parámetro:
//   - cadena: una cadena de texto para invertir.
//
// Retorna:
//   - La cadena invertida.
func reverseStringPro(cadena string) string {
	r := []rune(cadena) // Convierte la cadena a una lista de runas para soporte UTF-8.

	// Bucle que intercambia los caracteres de los extremos hacia el centro.
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
