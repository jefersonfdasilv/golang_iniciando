package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Programa principal.")
	auxiliar.Escrever("Olá, mundo!")
	emailValido := auxiliar.ValidarEmail("jeferson@rundev.com.br.")
	if !emailValido{
		fmt.Println("Email inválido")
	}
}
