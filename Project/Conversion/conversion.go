package conversion

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func StringsToFloat(stringss []string)([]float64,error){

	var floats []float64 

	for _,stringVal := range stringss{
		 
		stringVal := strings.TrimSpace(stringVal)
		floatPrice,err :=strconv.ParseFloat(stringVal,64)

		if (err!=nil){
			fmt.Print(err)
			return nil,errors.New("failed To Convert String to Float")
		} 

		floats = append(floats,floatPrice)
	}
	return floats,nil

}