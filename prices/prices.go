package prices

import (
	"bufio"
	"fmt"
	"os"
	"github.com/taylorjo02/price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadPrices()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       []float64{10, 20, 30},
		TaxIncludedPrices: map[string]float64{},
	}
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("unable to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("scanner error, %w", err)
		file.Close()
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
	}

	job.InputPrices = prices
	file.Close()
}
