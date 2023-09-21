package contas

type ContaCorrente struct {
	Titular string
	NumeroAgencia int
	NumeroConta int
	Saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque Realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if(valorDaTransferencia < c.Saldo && valorDaTransferencia > 0) {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if(valorDoDeposito > 0) {
		c.Saldo += valorDoDeposito
		return "Deposito Realizado com sucesso", c.Saldo
	}

	return "O valor do deposito e menor que 0,", c.Saldo
}