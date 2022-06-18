package contas

import "golang/clientes"

type ContaCorrente struct {
	Titular        clientes.Cliente
	Agencia, Conta int
	saldo          float64
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque >= 0 && conta.saldo >= valorDoSaque

	if !podeSacar {
		return "saldo insuficiente", conta.saldo
	}

	conta.saldo -= valorDoSaque
	return "Saque realizado com sucesso!", conta.saldo
}

func (conta *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {

	podeDepositar := valorDoDeposito > 0

	if !podeDepositar {
		return "Deposite um valor maior que zero", valorDoDeposito
	}

	conta.saldo += valorDoDeposito
	return "Depósito realizado com sucesso. Novo saldo", conta.saldo
}

func (conta *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {

	podeTransferir := valorDaTransferencia >= 0 && conta.saldo >= valorDaTransferencia
	if !podeTransferir {
		return "Não é possível realizar a transferência. Observe o valor informado e seu saldo na conta. saldo da conta: ", conta.saldo
	}

	conta.Sacar(valorDaTransferencia)
	contaDestino.Depositar(valorDaTransferencia)
	return "Transferência realizada com sucesso", contaDestino.saldo
}

func (conta *ContaCorrente) GetSaldo() (string, float64) {
	return "Seu saldo da conta corrente é de:", conta.saldo
}
