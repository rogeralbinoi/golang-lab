package main

import (
	"fmt"
	m "math" // criar alias para o pacote
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo inferido float64

	// forma reduzida
	area := PI * m.Pow(raio, 2)

	fmt.Printf("A área da circunferência é %.4f\n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var e, f bool = true, false

	g, h, i := 1, false, "epa"

	fmt.Println(a, b, c, d, e, f, g, h, i)

}
