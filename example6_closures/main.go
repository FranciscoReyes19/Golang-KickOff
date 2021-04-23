package main

import (
	"fmt"
)

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//funciones anonimas
	fmt.Printf("Suma de 4 + 5: %d \n", Calculo(5, 7))

	//resta
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resta de 4 + 5: %d \n", Calculo(15, 7))

	//division
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Division de 12 + 3: %d \n", Calculo(12, 3))

	Operaciones()

	/* CLOSURES */
	tabladDel := 2

	Mitabla := Table(tabladDel)

	for i := 1; i < 11; i++ {
		fmt.Println(Mitabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 14
		return a + b
	}
	fmt.Println(resultado())
}

//Closures
func Table(valor int) func() int {
	secuencia := 0
	num := valor

	return func() int {
		secuencia++
		return num * secuencia
	}
}
