package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	var prices []string
	fmt.Println("Please Enter your prices. Confirm every price with Enter")
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)

	}
	return prices, nil
}
func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}
func New() CMDManager {
	return CMDManager{}
}
