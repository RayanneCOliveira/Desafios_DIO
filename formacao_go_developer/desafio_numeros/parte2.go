/* 

Criar um código que imprima os números de 1 a 100 com as seguintes condições:
1- Exibir a palavra "Pin" sempre que aparecer um múltiplo de 3.
2- Exibir a palavra "Pan" sempre que aparecer um múltiplo de 5.

*/

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("PinPan")
		} else if i % 3 == 0 {
			fmt.Println("Pin")
		} else if i % 5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}