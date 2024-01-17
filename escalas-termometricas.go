package main

import "fmt"

const ebulicaoK = 273.0

func main() {
	var temperaturaK = 500.0
	var temperaturaC = temperaturaK - ebulicaoK
	fmt.Printf("A temperatura de ebulição em ºK é: %g e a temperatura em ºC é: %g", temperaturaK, temperaturaC)
}
