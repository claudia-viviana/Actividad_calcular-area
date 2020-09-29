// De la Cerda Rizo Claudia Viviana
// Calcular el área de un círculo

package main

import "fmt"

func main() {
	var radio float64
	const pi float64 = 3.141592

	fmt.Print("Radio: ")
	fmt.Scanf("%f", &radio)

	circulo := pi * radio * radio
	fmt.Println("Area del círculo: ", circulo)
}
