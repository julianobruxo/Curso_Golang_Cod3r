package main

import (
	"fmt"
)

// func refatoramento(password int) string {
// 	switch {
// 	case password > 60:
// 		return "Vai pra casa véio"
// 	case password < 60 && password >= 18:
// 		return "Tá de boas, pode entrar"
// 	case password < 18 && password > 5:
// 		return "Vaza pirralhada"
// 	default:
// 		return "Tá drogado?"
// 	}
// }

func recognition(user string) string {
	switch {
	case user == "Juliano":
		return "Welcome master. "
	case user == "Suzy":
		return "Welcome master's master."
	case user == "Isa":
		return "Welcome young master"
	default:
		return "Get lost modafaca"
	}
}

func main() {
	fmt.Println("Welcome")
	var name string
	fmt.Println("Type your name")
	fmt.Scan(&name)
	fmt.Println(recognition(name))
	var p int
	fmt.Println("Type your password")
	fmt.Scan(&p)
	if p == 12345 {
		fmt.Println("Access Granted. You may proceed.")
	} else {
		fmt.Println("Access Denied. Terminating program. \nTry again")
	}

	// if access(p) == true {
	// 	return "Access granted."
	// } else {
	// 	return "Wrong password. Try Again"
	// }

}
