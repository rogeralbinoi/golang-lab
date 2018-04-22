package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f.", x)

	a, b, c, d := 1, 1.9999, false, "Opa!"

	fmt.Printf("\na: %d, b: %.4f, c: %t, d: %s\n", a, b, c, d)
	fmt.Printf("\na: %v, b: %v, c: %v, d: %v\n", a, b, c, d)

}
