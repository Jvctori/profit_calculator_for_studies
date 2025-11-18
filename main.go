package main

import (
	"fmt"
)

func main() {
	var rv float64
	var exp float64
	var tr float64
	// não adiciona quebra de linha no final, como fmt.Println()
	//fmt.Print("Digite sua receita: ")
	// fetching data from user / capturando dados do usuário
	//fmt.Scan(&revenue)
	fetchValue("Revenue: ", &rv)

	/*fmt.Print("Digite suas despesas: ")
	fmt.Scan(&expenses)*/
	fetchValue("Expenses: ", &exp)

	/*fmt.Print("Digite sua a '%' de imposto: ")
	fmt.Scan(&tax_rate)*/
	fetchValue("Tax Rate: ", &tr)

	ebt, prf, rt := calculateProfit(rv, exp, tr)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", prf)
	fmt.Printf("Ratio: %.2f\n", rt)
	/*ebt := revenue - expenses
	tax_payed := (ebt * (tax_rate / 100))
	ratio := ebt / (ebt - tax_payed)*/

	/*fmt.Print("Ganhos antes dos impostos: ")
	fmt.Println(ebt)

	fmt.Print("Lucro depois dos impostos: ")
	fmt.Println(ebt - tax_payed)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)*/
}

/* Função que calcula o lucro*/
func calculateProfit(r, e, t float64) (float64, float64, float64) {
	ebt := r - e
	p := ebt * (1 - t/100)
	rt := ebt / p

	return ebt, p, rt
}

/*
Função que imprime um texto e captura um valor float64 do usuário
*float64 representa um ponteiro (endereço na memória):
permitindo alterar o valor diretamente na memória
*/
func fetchValue(text string, v *float64) {
	fmt.Print(text)
	fmt.Scan(v)
}

