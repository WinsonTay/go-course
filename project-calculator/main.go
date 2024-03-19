package main

import (
	"fmt"

	//"example.com/calculator/prices"
	"example.com/calculator/filemanager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", taxRate*100))
		//cmdManager := cmdmanager.New()
		taxIncludedprices := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go taxIncludedprices.Process(doneChans[i], errorChans[i])
	}
	// Wait for all processes to finish and check if there were any errors
	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err {
				fmt.Printf("Error in process with tax rate: %.2f%%\n", taxRates[index]*100) // print the tax rate that caused the error
			}
		case <-doneChans[index]:
			fmt.Printf("Done!\n")
		}
	}

}
