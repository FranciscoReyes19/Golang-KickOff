package main

import (
	"fmt"
	"strconv"
)

var numero int
var Texto string = "Soy exportable"
var status bool = true

func main() {
	numero2, numero3, status, texto := 2, 4, false, "Otro"

	var floatNum int64 = 6

	numero2 = int(floatNum)
	//texto = strconv.FormatInt(int64(numero2), 10)
	texto = strconv.Itoa(int(numero2))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(status)
	fmt.Println(texto)

	usingScope()
}

func usingScope() {
	fmt.Println(status)
}
