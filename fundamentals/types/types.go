package main

import (
	"fmt"
	m "math"
	r "reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println(r.TypeOf(1), r.TypeOf(2), r.TypeOf(1000))

	// sem sinal (só positivo)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println(r.TypeOf(b))

	// com sinal
	i1 := m.MaxInt64
	fmt.Println("Valor máximo de int é", i1, r.TypeOf(i1))

	const i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println(i2, r.TypeOf(i2))

	// Números reais(float32, float64)
	var x = float32(30.9) // float64 por padrao
	fmt.Println(x, r.TypeOf(x))

	// boolean
	bo := false

	fmt.Println(bo, r.TypeOf(bo))

	// string

	s1 := "String aqui"

	fmt.Println(s1 + "!")

	// string com multiplas linhas

	s2 := `
	  1: primeira linha
	  2: segunda linha
	`

	fmt.Println(s2)

	// char?? não existe!
	// chat := 'a'

}
