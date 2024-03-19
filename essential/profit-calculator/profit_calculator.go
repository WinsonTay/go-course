package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type BalanceSheet struct {
	EarningsBeforeTax float64 `json:"earningsBeforeTax"`
	EarningsAfterTax  float64 `json:"earningsAfterTax"`
	Ratio             float64 `json:"ratio"`
}

func writeBalanceSheetToJson(balanceSheet BalanceSheet) {

	balanceSheetJson, _ := json.MarshalIndent(balanceSheet, "", "")
	os.WriteFile("balance.json", []byte(balanceSheetJson), 0644)
}
func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var err error
	revenue, expenses, taxRate, err = getUserInput()
	if err != nil {
		fmt.Println(err)
		panic("Exiting Program ..")
	}
	earningsBeforeTax, earningAfterTax, ratio := calculateProfits(revenue, expenses, taxRate)
	balanceSheet := BalanceSheet{
		EarningsBeforeTax: earningsBeforeTax,
		EarningsAfterTax:  earningAfterTax,
		Ratio:             ratio,
	}

	writeBalanceSheetToJson(balanceSheet)

	fmt.Printf("EBT: %.1f ", earningsBeforeTax)
	fmt.Printf("Profit: %.1f ", earningAfterTax)
	fmt.Printf("Ratio : %.1f ", ratio)

}

func getUserInput() (rev float64, exp float64, tax float64, err error) {
	fmt.Print("What is the Revenue: ")
	fmt.Scan(&rev)

	if rev < 0 {
		err = errors.New("Revenue cannot or negative number")
	}
	fmt.Print("What is the Expenses: ")
	fmt.Scan(&exp)

	if exp < 0 {
		err = errors.New("Expenses cannot or negative number")
	}
	fmt.Print("Tax Rate: ")
	fmt.Scan(&tax)

	if tax < 0 {
		err = errors.New("Tax cannot or negative number")
	}
	return rev, exp, tax, err
}

func calculateProfits(revenue float64, expenses float64, taxRate float64) (ebt float64, eft float64, ratio float64) {
	ebt = revenue - expenses
	eft = ebt * (1 - taxRate/100)
	ratio = ebt / eft
	return ebt, eft, ratio
}
