package main

import (
	"fmt"
)

func main() {
	exponencia(2)
}

func exponencia(num int) {
	if num > 100000 {
		return
	}
	fmt.Println(num)
	exponencia(num * 2)
}
