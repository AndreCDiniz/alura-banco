package contas

import "github.com/AndreCDiniz/banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

//c *ContaCorrente funciona como se fosse o this para apontarmos direto para a conta corrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeDeposito float64) (string, float64) {
	if valorDeDeposito > 0 {
		c.saldo += valorDeDeposito

		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

// func entendendoPonteiros() {
// 	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroConta: 1252, numeroAgencia: 3295, saldo: 123.6}
// 	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", numeroConta: 1252, numeroAgencia: 3295, saldo: 123.6}

// 	//Aqui embaixo é outra forma que podemos instanciar uma struct
// 	// contaDoAndre := ContaCorrente{"Andre", 1483, 3295, 294.6}
// 	// fmt.Println(contaDoAndre)

// 	fmt.Println(contaDoGuilherme, contaDoGuilherme2)
// 	fmt.Println(contaDoGuilherme == contaDoGuilherme2)

// 	fmt.Println()
// 	var contaDaCris *ContaCorrente
// 	contaDaCris = new(ContaCorrente)
// 	contaDaCris.titular = "Cris"
// 	contaDaCris.numeroConta = 500

// 	var contaDaCris2 *ContaCorrente
// 	contaDaCris2 = new(ContaCorrente)
// 	contaDaCris2.titular = "Cris"
// 	contaDaCris2.numeroConta = 500

// 	fmt.Println(contaDaCris, contaDaCris2)
// 	fmt.Println(&contaDaCris, &contaDaCris2)
// 	//Aqui está dando falso pois estamos comparando o endereco de memoria
// 	fmt.Println(contaDaCris == contaDaCris2)
// 	fmt.Println()
// 	fmt.Println(*contaDaCris, *contaDaCris2)
// 	//Aqui está dando true pois estamos comparando o conteudo
// 	fmt.Println(*contaDaCris == *contaDaCris2)
// }
