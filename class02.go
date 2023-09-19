package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
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

func main() {
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)

	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500
	fmt.Println(contaDaCris)
	fmt.Println(*contaDaCris)

	contaDoGuilhere := ContaCorrente {titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDoGuilhere2 := ContaCorrente {titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}

	fmt.Println(contaDoGuilhere == contaDoGuilhere2)

	contaDoGuilhere3 := ContaCorrente {titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDoGuilhere4 := ContaCorrente {titular: "Guilherme", numeroAgencia: 500, numeroConta: 123456, saldo: 125.50}

	fmt.Println(contaDoGuilhere3 == contaDoGuilhere4)
	fmt.Println(&contaDoGuilhere3)
	fmt.Println(&contaDoGuilhere4)
	fmt.Println(&contaDoGuilhere4)

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)

	var contaDaCris3 *ContaCorrente
	contaDaCris3 = new(ContaCorrente)

	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500
	contaDaCris3.titular = "Cris"
	contaDaCris3.saldo = 500

	fmt.Println(contaDaCris2 == contaDaCris3)
	fmt.Println(contaDaCris2 == contaDaCris3)

	contaDaSilva := ContaCorrente{}
	contaDaSilva.titular = "Silva"
	contaDaSilva.saldo = 500

	fmt.Println(contaDaSilva.saldo)
	valorDoSaque := 200.
	contaDaSilva.saldo = contaDaSilva.saldo - valorDoSaque

	fmt.Println(contaDaSilva.saldo)

	valorDoSaque = 800.
	contaDaSilva.saldo = contaDaSilva.saldo - valorDoSaque
	fmt.Println(contaDaSilva.saldo)

	contaDaSilva2 := ContaCorrente{}
	contaDaSilva2.titular = "Silva"
	contaDaSilva2.saldo = 500

	valorDoSaque2 := 200.
	fmt.Println(contaDaSilva2.saldo)
	fmt.Println(contaDaSilva2.Sacar(valorDoSaque2))
	
	valorDoSaque3 := 800.
	fmt.Println(contaDaSilva2.Sacar(valorDoSaque3))

}