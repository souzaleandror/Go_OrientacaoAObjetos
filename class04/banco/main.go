package main

import (
	"fmt"
	
	"example.com/module/Users/newvoicedevelopment/Documents/banco/contas"
	"example.com/module/Users/newvoicedevelopment/Documents/banco/clientes"
)

func main() {
	contaDaSilva := contas.ContaCorrente{}
	contaDaSilva.Titular = "Silva"
	contaDaSilva.Saldo = 500

	fmt.Println(contaDaSilva.Saldo)
	fmt.Println(contaDaSilva.Depositar(200))
	fmt.Println(contaDaSilva.Saldo)

	status, valor := contaDaSilva.Depositar(200)
	fmt.Println(status)
	fmt.Println(valor)

	contaDaLaura := contas.ContaCorrente{Titular: "Laura", Saldo: 300}
	contaDaGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 500}

	fmt.Println(contaDaLaura)
	fmt.Println(contaDaGustavo)

	status2 := contaDaLaura.Transferir(-200, &contaDaGustavo)
	fmt.Println(status2)

	fmt.Println(contaDaLaura)
	fmt.Println(contaDaGustavo)

}