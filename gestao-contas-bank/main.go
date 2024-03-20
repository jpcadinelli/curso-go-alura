package main

import (
	"curso-go-alura/gestao-contas-bank/contas"
	"fmt"
)

func main() {
	contaDoJoao := contas.ContaCorrente{Titular: "João Pedro", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 158.99}
	fmt.Println(contaDoJoao.Saldo)
	contaDaBea := contas.ContaCorrente{Titular: "Beatrice", NumeroAgencia: 589, NumeroConta: 654321, Saldo: 12158.99}
	fmt.Println(contaDaBea.Saldo)
	status := contaDaBea.Tranferir(545.02, &contaDoJoao)
	fmt.Println(status)
	fmt.Println(contaDoJoao.Saldo)
	fmt.Println(contaDaBea.Saldo)
}
