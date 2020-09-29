// De la Cerda Rizo Claudia Viviana
// Calcular el área de un cuadrado

package main

import "fmt"

func main() {
	var lado float64

	fmt.Print("Lado: ")
	fmt.Scanf("%f", &lado)

	cuadrado := lado * lado
	fmt.Println("Area del cuadrado: ", cuadrado)
}
