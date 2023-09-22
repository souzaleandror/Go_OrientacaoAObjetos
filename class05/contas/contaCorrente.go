package contas

	import 	"class05/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta  int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque Realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if(valorDaTransferencia < c.saldo && valorDaTransferencia > 0) {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if(valorDoDeposito > 0) {
		c.saldo += valorDoDeposito
		return "Deposito Realizado com sucesso", c.saldo
	}

	return "O valor do deposito e menor que 0,", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}