package main

import "fmt"

func main() {
	fmt.Println(Calculo(9))
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(5, 4, 3))
	fmt.Println(Calculo(5, 46, 44))
	fmt.Println(Calculo(5, 46, 1, 45))

}

func uno(numero int) (m int) {
	m = numero * 2
	return
}

func dos(num int) (int, bool) {
	if num == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

//parametros variables
// _  -> significa que no se requiere una variable para el iterable (0 1 2)
func Calculo(num ...int) int {
	total := 0
	for _, n := range num {
		total = total + n
	}
	return total
}
