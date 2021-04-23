package main

import (
	"fmt"
)

//var table = [10]int{5, 8, 7, 9, 7, 4, 8, 12, 9, 3} //global

func main() {
	//table := [10]int{5, 8, 7, 9, 7, 4, 8, 12, 9, 3} //local

	//slice -> vectores dianamicos
	matriz := []int{3, 7, 54}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[:4]

	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo: %d, Capacidad: %d", len(elementos), cap(elementos))

}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 130; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n Largo: %d, Capacidad: %d", len(nums), cap(nums))

}
