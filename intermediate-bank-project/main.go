package main

import "fmt"

type ContaCorrente struct {
	titular     string
	agencia     int
	numeroConta int
	saldo       float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com sucesso", c.saldo
	}

	return "Saldo insuficiente", c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

	return "valor do deposito menor do que zero", c.saldo

}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500.

	fmt.Println(contaDaSilvia.saldo)

	status, valor := contaDaSilvia.Depositar(2000.)
	fmt.Println(status, valor)
}
