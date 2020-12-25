package auxiliar

import (
	"fmt"
	"github.com/badoux/checkmail"
)
//Função para escrever uma string na saída padrão
func Escrever(msg string) {
	fmt.Println(msg)
	escrever("Função privada.")
}
//Função para validar se dada uma string representa um formato válido de e-mail
func ValidarEmail(email string) bool  {
	err := checkmail.ValidateFormat(email)
	return err == nil
}