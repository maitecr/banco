package main

import (
	"fmt"

	contas "banco/conta"
)

func main() {

	conta01 := contas.ContaCorrente{
		Titular:       "Pafuncio",
		NumeroAgencia: 123,
		NumeroConta:   01,
	}
	fmt.Println(conta01)

	conta02 := contas.ContaCorrente{
		Titular:       "Pafuncio",
		NumeroAgencia: 123,
		NumeroConta:   01,
	}
	fmt.Println(conta02)

	//fmt.Println(conta01 == conta02) //mesmo que os endereços sejam diferentes, o go consegue identificar o conteúdo do objeto

	//conta02 := ContaCorrente{"Marisa", 123, 02, 100.40} //usar para ter todos os campos preenchidos
	//fmt.Println(conta02)

	var conta03 *contas.ContaCorrente //ponteiro para a struct ContaCorrente,o print retornará com um &
	conta03 = new(contas.ContaCorrente)
	conta03.Titular = "Aristoteles"
	//fmt.Println(&conta03) //usar o * para obter o conteúdo que há no ponteiro/endereço de memória

	var conta04 *contas.ContaCorrente
	conta04 = new(contas.ContaCorrente)
	conta04.Titular = "Aristoteles"
	conta04.Saldo = 400
	//fmt.Println(&conta04) //usar o * para obter o conteúdo que há no ponteiro/endereço de memória

	//fmt.Println(conta03 == conta04) //endereços diferentes = false

	fmt.Println("===============================================================================================")

	conta05 := contas.ContaCorrente{Titular: "Penelope", Saldo: 500}
	fmt.Println(conta05)

	conta06 := contas.ContaCorrente{Titular: "Penelope", Saldo: 500}

	//status, valor := conta05.Depositar(2000)
	//fmt.Println(status, valor)

	//fmt.Println(conta05.Sacar(-100))
	//fmt.Println(conta05.saldo)
	//fmt.Println(conta05.Depositar(400))
	//fmt.Println(conta05.saldo)

	status := conta05.Transferir(100, &conta06)

	fmt.Println(status)
	fmt.Println(conta05)
	fmt.Println(conta06)
}
