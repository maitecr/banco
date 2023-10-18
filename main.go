package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso. Saldo:", c.saldo
	} else {
		return "Valor inválido para depósito Saldo:", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)

		return true
	} else {
		return false
	}
}

func main() {

	conta01 := ContaCorrente{
		titular:       "Pafuncio",
		numeroAgencia: 123,
		numeroConta:   01,
	}
	fmt.Println(conta01)

	conta02 := ContaCorrente{
		titular:       "Pafuncio",
		numeroAgencia: 123,
		numeroConta:   01,
	}
	fmt.Println(conta02)

	//fmt.Println(conta01 == conta02) //mesmo que os endereços sejam diferentes, o go consegue identificar o conteúdo do objeto

	//conta02 := ContaCorrente{"Marisa", 123, 02, 100.40} //usar para ter todos os campos preenchidos
	//fmt.Println(conta02)

	var conta03 *ContaCorrente //ponteiro para a struct ContaCorrente,o print retornará com um &
	conta03 = new(ContaCorrente)
	conta03.titular = "Aristoteles"
	//fmt.Println(&conta03) //usar o * para obter o conteúdo que há no ponteiro/endereço de memória

	var conta04 *ContaCorrente
	conta04 = new(ContaCorrente)
	conta04.titular = "Aristoteles"
	conta04.saldo = 400
	//fmt.Println(&conta04) //usar o * para obter o conteúdo que há no ponteiro/endereço de memória

	//fmt.Println(conta03 == conta04) //endereços diferentes = false

	fmt.Println("===============================================================================================")

	conta05 := ContaCorrente{titular: "Penelope", saldo: 500}
	fmt.Println(conta05)

	conta06 := ContaCorrente{titular: "Penelope", saldo: 500}

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
