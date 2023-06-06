package main

import "fmt"

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

func main() {
	pontoEbulicaoC := 100.0
	pontoEbulicaoK := celsiusToKelvin(pontoEbulicaoC)

	fmt.Printf("O ponto de ebulição da água em Celsius é %.2f °C.\n", pontoEbulicaoC)
	fmt.Printf("O ponto de ebulição da água em Kelvin é %.2f K.\n", pontoEbulicaoK)
}
