package main

//Soma com for normal, mas podemos usar o range, para simplificar o codigo

// func Soma(numeros [5]int) int {
// 	soma := 0
// 	for i := 0; i < 5; i++ {
// 		soma += numeros[i]
// 	}
// 	return soma
// }

//Utilizando o range: ele vai executar de acordo com o tamanho do array, sem precisar repetir isso

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

//Aqui usamos o make para criar a slice, funciona, mas podemoa refatorar utilizando o append

// func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
// 	quantidadeDeNumeros := len(numerosParaSomar)
// 	somas = make([]int, quantidadeDeNumeros)

// 	for i, numeros := range numerosParaSomar {
// 		somas[i] = Soma(numeros)
// 	}

// 	return
// }

//Utilizando o append conseguimos deixar nosso for mais simples, juntando duas slices

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}

	return somas
}
