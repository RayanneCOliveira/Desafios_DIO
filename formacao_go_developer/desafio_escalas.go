/* Desenvolver um programa para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius

C = K - 273

*/

package main

import "fmt"

const tempK = 373

func main() {
	tempC := tempK - 273
	fmt.Printf("O ponto de ebulição da água em graus Kelvin é igual a %d e em graus Celsius é igual a %d.", tempK, tempC)
}
