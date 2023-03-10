package carteira

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

// O * antes do nome Carteira, indica que estamos trabalhando com ponteiro
func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

// Função para coloar BTC quando utilizar o %s no Errorf
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Usado para ver que o endereço é difirente no teste e no metodo e assim, utilizar ponteiro para resolver
// func (c Carteira) Depositar(quantidade int) {
// 	fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
// 	c.saldo += quantidade
// }
