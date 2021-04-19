package main

import "fmt"

var state bool

func main() {
	/* state = true
	if state = false; state {
		fmt.Println(state)
	} else {
		fmt.Println("false")
	}*/

	//switch

	switch num := 2; num {
	case 1:
		{
			fmt.Println("es 1")
		}
	case 2:
		{
			fmt.Println("es 2")
		}
	case 3:
		{
			fmt.Println("es 3")
		}
	case 4:
		{
			fmt.Println("es 4")
		}
	default:
		{
			fmt.Println("default")
		}

	}

}
