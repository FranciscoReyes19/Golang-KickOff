package main

/*
import (
	"bufio"
	"fmt"
	"os"
)
*/

var num1 int
var num2 int
var leyenda string
var resultado int

func main() {

	/*
		//using NewScanner
		fmt.Println("Ingrese numero 1: ")
		fmt.Scanln(&num1)

		fmt.Println("Ingrese numero 2: ")
		fmt.Scanln(&num2)

		fmt.Println("Descripcion : ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			leyenda = scanner.Text()
		}
		resultado = num1 + num2
		fmt.Println(leyenda, resultado)
	*/

	/*
		//using break
			num := 0

			for {
				fmt.Println("Continuo")
				fmt.Println("Ingrese el numero secreto")
				fmt.Scanln(&num)

				if num == 100 {
					break
				}
			}*/

	/*
		//using continue
			var i = 0

			for i < 10 {
				fmt.Printf("\n valor de i: %d", i)
				if i == 5 {
					fmt.Printf(" multiplicamos por 2 \n ")
					i = i * 2
					continue   //para volver a chquear la condicion, regresar al inicio del for
				}
				fmt.Printf("  Paso por aqui \n")
				i++
			}*/

	//using goto
	/*
		i := 0
		RUTINA:
			for i <= 10 {
				if i == 4 {
					i = i + 2
					fmt.Println("Voy a rutina suamando 2 a i")
					goto RUTINA //va ir a donde este el nombre RUTINA
				}
				fmt.Printf("valor de i: %d \n", i)
				i++
			}
	*/

}
