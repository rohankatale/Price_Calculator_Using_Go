package main

import (
	"fmt"

	//"example.com/price_calc/cmdmanager"
	"example.com/price_calc/filemanager"
	"example.com/price_calc/prices"
)


func main() {
	taxRates := []float64{0.1, 0.07, 0.33, 0.24, 0.15}

	doneChans := make([]chan bool,len(taxRates))
	errorChans := make([]chan error,len(taxRates))


	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error) 

		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))

		//cmdm := cmdmanager.New()
		

		priceJob:=prices.NewTaxIncludedPriceJob(fm,taxRate)
		//priceJob:=prices.NewTaxIncludedPriceJob(cmdm,taxRate)
		go priceJob.Process(doneChans[i],errorChans[i]) 
	
		// if err!=nil{
		// 	fmt.Println("could not process the job")
		// 	fmt.Println(err)
		// }
	
	}

	for i := range taxRates{
		select{
		case err:= <-errorChans[i]:
			if err!=nil{
				fmt.Println(err)
			}
		case  <-doneChans[i]:
			fmt.Println("done")
		}
	}


} 