/*
Archivo: main.go
Autor: Edwin Yoner
Fecha: 10-11-2024
Versión: 1.0

Descripción:
Este programa encuentra el número más grande en un slice de enteros utilizando la función `findLargest2`.
Si el slice está vacío, la función retorna un error. El programa prueba la función con distintos casos,
incluyendo números positivos, negativos y un slice vacío.
*/

package main

import (
	"errors" // Paquete para manejar errores en Go
	"fmt"    // Paquete para imprimir mensajes en la consola
)

// findLargest2 encuentra el número más grande en un slice de enteros.
// Parámetro:
//   - nums: un slice de enteros a evaluar.
//
// Retorna:
//   - El número más grande encontrado en el slice.
//   - Un error si el slice está vacío.
func findLargest2(nums []int) (int, error) {
	// Verificar si el slice está vacío para evitar errores al acceder a elementos
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

	max := nums[0] // Inicializar el máximo con el primer elemento

	// Iterar sobre el slice a partir del segundo elemento
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i] // Actualizar el máximo si se encuentra un número mayor
		}
	}

	return max, nil // Retornar el número más grande encontrado y nil como error
}

// Función principal del programa. Prueba la función findLargest2 con distintos casos de entrada.
func main() {
	fmt.Println(findLargest2([]int{3, 5, 7, 6, 2, 8, 9}))        // Caso con números positivos
	fmt.Println(findLargest2([]int{-2, -5, -7, -9, -3, -8, -6})) // Caso con números negativos
	fmt.Println(findLargest2([]int{}))                           // Caso con slice vacío
}
