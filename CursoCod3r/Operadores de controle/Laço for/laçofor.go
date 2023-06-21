package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1        /* ==> similar ao while*/
	for i <= 15 { /* não existe while em GO*/
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { /* for clássico*/
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ { /* for com estrutura bool */
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	for {
		/*laço infinito*/
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 2)
	}

}
