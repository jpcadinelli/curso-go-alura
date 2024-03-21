package contas

import "gestaoContasBank/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso.", c.Saldo
	} else {
		return "Saldo insuficiente.", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.Saldo
	} else {
		return "Valor negativo não pode ser depositado.", c.Saldo
	}
}

func (c *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
