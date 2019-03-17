package main

import (
	"fmt"

	"github.com/albarin/indexa/pkg/indexa"
)

const indexaURL = "https://api.indexacapital.com"
const indexaToken = "foo"

func main() {
	c := indexa.NewIndexaClient(indexaURL, indexaToken)

	me, err := c.Me()
	if err != nil {
		fmt.Println("error:", err)
	}

	p, err := c.Performance(me.Accounts[0].AccountNumber)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Invertido: %.2f€\n", float64(p.Return.Investment))
	fmt.Printf("Rentabilidad: %.2f€\n", p.Return.Pl)
	fmt.Printf("Total: %.2f€\n", p.Return.TotalAmount)
	fmt.Printf("Rentabilidad ponderada por tiempo: %.1f%% acumulada (%.1f%% TAE)\n", p.Return.TimeReturn*100, p.Return.TimeReturnAnnual*100)
	fmt.Printf("Rentabilidad ponderada por dinero: %.1f%% acumulada (%.1f%% TAE)\n", p.Return.MoneyReturn*100, p.Return.MoneyReturnAnnual*100)
}
