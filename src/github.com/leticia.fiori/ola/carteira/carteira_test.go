package carteira

import "testing"

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	//Utilizando o %s ele tras o BTC criado no metodo String, se utilizar %d, não tras o BTC
	if resultado != esperado {
		t.Errorf("resultado %s, esperado %s", resultado, esperado)
	}
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

func confirmaErro(t *testing.T, resultado error, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu")
	}

	//Utilizando o %s ele tras o BTC criado no metodo String, se utilizar %d, não tras o BTC
	if resultado != esperado {
		t.Errorf("erro resultado %s, erro esperado %s", resultado, esperado)
	}
}

// //Usado para ver que o endereço é difirente no teste e no metodo e assim, utilizar ponteiro para resolver
// func TestCarteira(t *testing.T) {
// 	carteira := Carteira{}

// 	carteira.Depositar(10)

// 	resultado := carteira.Saldo()

// 	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

// 	esperado := 10

// 	if resultado != esperado {
// 		t.Errorf("resultado %d, esperado %d", resultado, esperado)
// 	}
// }
