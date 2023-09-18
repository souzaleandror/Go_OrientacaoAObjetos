package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main() {
	var titular string = "Guilherme"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50

	var titular2 string = "Luciene"

	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	fmt.Println(titular2)

	fmt.Println(ContaCorrente{})

	var contaDoGuilhere ContaCorrente = ContaCorrente {}
	contaDoGuilhere2 := ContaCorrente {}
	contaDoGuilhere3 := ContaCorrente {titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDoGuilhere4 := ContaCorrente {titular: "Guilherme", saldo: 125.50}
	contaDaBruna := ContaCorrente {"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilhere)
	fmt.Println(contaDoGuilhere2)
	fmt.Println(contaDoGuilhere3)
	fmt.Println(contaDoGuilhere4)
	fmt.Println(contaDaBruna)
}