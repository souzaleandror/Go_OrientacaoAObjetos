package contas

import 	"class05/clientes"

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao  int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
	if(valorDaTransferencia < c.saldo && valorDaTransferencia > 0) {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if(valorDoDeposito > 0) {
		c.saldo += valorDoDeposito
		return "Deposito Realizado com sucesso", c.saldo
	}

	return "O valor do deposito e menor que 0,", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}