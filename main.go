package main

import (
	"fmt"
	"golang/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {

	contaPoupanca := contas.ContaPoupanca{}
	contaCorrente := contas.ContaCorrente{}

	contaPoupanca.Depositar(200.0)
	contaCorrente.Depositar(100.0)

	PagarBoleto(&contaCorrente, 60.0)
	PagarBoleto(&contaPoupanca, 100.0)
	fmt.Println(contaCorrente.GetSaldo())
	fmt.Println(contaPoupanca.GetSaldo())
}
