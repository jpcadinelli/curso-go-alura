package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}


func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso.", c.saldo
	} else {
		return "Saldo insuficiente.", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso.", c.saldo
	} else {
		return "Valor negativo não pode ser depositado.", c.saldo
	}
}

func (c *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDoJoao := ContaCorrente{titular: "João Pedro", numeroAgencia: 589, numeroConta: 123456, saldo: 158.99}
	fmt.Println(contaDoJoao.saldo)
	contaDaBea := ContaCorrente{titular: "Beatrice", numeroAgencia: 589, numeroConta: 654321, saldo: 12158.99}
	fmt.Println(contaDaBea.saldo)
	status := contaDaBea.Tranferir(545.02, &contaDoJoao)
	fmt.Println(status)
	fmt.Println(contaDoJoao.saldo)
	fmt.Println(contaDaBea.saldo)
}
