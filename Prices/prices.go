package prices

import (
	"fmt"

	"example.com/price_calc/conversion"
	"example.com/price_calc/iomanager"
)

type TaxIncludedPriceJob struct{
	IOManager  iomanager.IOManager      `json:"-"`
	TaxRate float64						`json:"tax_rate"`
	InputPrices []float64				`json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}
 

func (job *TaxIncludedPriceJob) Process() error{
	err := job.LoadData() 
	if err!=nil{
		return err
	}
	result := make(map[string]string)
		for _, price := range job.InputPrices {
			taxIncludedPrice:= price + price*job.TaxRate
			result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxIncludedPrice)
		}
	fmt.Println(result)
	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		IOManager: fm,
		InputPrices: []float64{10, 20, 30, 40, 50},
		TaxRate : taxRate,
	}
}


func (job *TaxIncludedPriceJob) LoadData() error{

	data,err:= job.IOManager.ReadLines()
	if err != nil{
		return err
	}
	prices,err := conversion.StringsToFloat(data)
	if err!=nil{
		return err
	}
	job.InputPrices = prices
	return nil
}