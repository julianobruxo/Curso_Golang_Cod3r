package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 1 && t.Hour() < 13:
		fmt.Println("Bom dia")
	case t.Hour() > 12 && t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")

	}
}
