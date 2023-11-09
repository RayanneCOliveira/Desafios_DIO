// Criar um código que exiba todos os números entre 1 e 100 divisíveis por 3.

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}