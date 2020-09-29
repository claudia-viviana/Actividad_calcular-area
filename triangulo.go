// De la Cerda Rizo Claudia Viviana
// Calcular el área de un triángulo

package main

import "fmt"

func main() {
	var base float64
	var altura float64

	fmt.Print("Base: ")
	fmt.Scan(&base)

	fmt.Printf("Altura: ")
	fmt.Scan(&altura)

	triangulo := (base * altura) / 2
	fmt.Println("Area del triángulo: ", triangulo)
}
