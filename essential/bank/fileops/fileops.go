package fileops

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type BalanceSheet struct {
	Balance float64 `json:"balance"`
}

func GetBalanceFromFile(fileName string) (float64, error) {
	balanceSheet := BalanceSheet{}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to Get File")
	}
	err = json.Unmarshal(data, &balanceSheet)

	if err != nil {
		errMsg := fmt.Sprintf("Error Parsed to Json %v", err.Error())
		return 1000, errors.New(errMsg)
	}
	return balanceSheet.Balance, nil
}
func WriteBalanceToFile(balance float64) {
	balanceSheet := BalanceSheet{
		Balance: balance,
	}
	balanceSheetJson, _ := json.MarshalIndent(balanceSheet, "", "")
	os.WriteFile("balance.json", []byte(balanceSheetJson), 0644)

}
