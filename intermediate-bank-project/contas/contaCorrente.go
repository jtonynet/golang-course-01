package contas

import "github.com/jtonynet/golang-course-01/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com sucesso", c.saldo
	}

	return "saldo insuficiente", c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	}

	return "valor do deposito menor do que zero", c.saldo

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)

		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
