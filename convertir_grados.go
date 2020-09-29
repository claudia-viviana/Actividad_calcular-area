// De la Cerda Rizo Claudia Viviana
// Convertir de grados Fahrenheit a Celsius (C = (F − 32) * 5/9)

package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Ingresa los grados en Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("Grados en Celsius: ", celsius)
}
