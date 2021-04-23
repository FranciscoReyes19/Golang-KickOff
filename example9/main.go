package main

import (
	"fmt"
	"time"

	us "./user"
)

type fran struct {
	us.Usuario
}

func main() {
	u := new(fran)
	u.Signup(1, "Fran", time.Now(), true)
	fmt.Println(u.Usuario)
}
