package main

import (
	"fmt"

	contas "github.com/jtonynet/golang-course-01/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	// Transferir
	/*
		contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300.}
		contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100.}

		status := contaDaSilvia.Transferir(-200, &contaDoGustavo)

		fmt.Println(status)
		fmt.Println(contaDaSilvia)
		fmt.Println(contaDoGustavo)
	*/

	// STRUCTS Aninhadas
	/*
		clienteBruno := clientes.Titular{
			Nome:      "Bruno",
			CPF:       "000.000.000.00",
			Profissao: "Desenvolvedor",
		}

		contaDoBruno := contas.ContaCorrente{
			Titular:       clienteBruno,
			NumeroAgencia: 123,
			NumeroConta:   123456,
		}

		contaDoBruno.Depositar(500)

		fmt.Println(contaDoBruno)
		contaDoBruno.Depositar(100)

		fmt.Println(contaDoBruno.ObterSaldo())
	*/

	// Interfaces
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())
}
