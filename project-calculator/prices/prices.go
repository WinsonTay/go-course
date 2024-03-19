package prices

import (
	"errors"
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"taxRate"`
	InputPrices       []float64           `json:"inputPrices"`
	TaxIncludedPrices map[string]string   `json:"taxIncludedPrices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
func (job *TaxIncludedPriceJob) LoadData() error {

	var prices []float64
	var err error
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to read lines")
	}
	prices, err = conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return errors.New("failed to parse")
	}
	job.InputPrices = prices
	return nil

}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan bool) {
	// Get Prices
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxedPrice := price * (1 + job.TaxRate)
		formattedPrice := fmt.Sprintf("%.2f", taxedPrice)
		result[fmt.Sprintf("%.2f", price)] = formattedPrice
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)

	doneChan <- true
}
