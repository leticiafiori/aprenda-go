package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const ingles = "ingles"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaIngles = "Hello, "

func main() {
	fmt.Println(Ola("mundo", ""))
}

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
	//retonando com if

	/* 	if idioma == espanhol {
	   		return prefixoOlaEspanhol + nome
	   	}

	   	if idioma == frances {
	   		return prefixoOlaFrances + nome
	   	}

	   	return prefixoOlaPortugues + nome */

}

func prefixoSaudacao(idioma string) (prefixo string) {
	//retornando com switch case
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
