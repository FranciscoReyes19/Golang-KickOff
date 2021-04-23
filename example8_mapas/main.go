package main

import (
	"fmt"
)

func main() {
	paises := make(map[string]string, 5)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos aires"

	campeonato := map[string]int{
		"Aola":  5,
		"Rola2": 6,
		"Cola3": 7,
		"Dola4": 17,
	}

	campeonato["Mola5"] = 67

	delete(campeonato, "Hola2")

	for equipo, score := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de %d \n", equipo, score)
	}

	score, existe := campeonato["Dola4"]

	fmt.Printf("El puntaje campturado es %d y el equipo existe %t \n", score, existe)

}
