package main

import "fmt"

var itemsTax = map[float64]float64{
	0.99: 0.75,
	2.75: 0.15,
	0.87: 0.2,
}

func main() {

	totalSalesTax:=getTotalSalesTax(itemsTax)

	fmt.Println("Sales Tax Total: ",totalSalesTax)
}

func calculateSalesTax(item, taxRate float64) float64 {
	return item*taxRate
}

func getTotalSalesTax(record map[float64]float64) float64 {

	var salesTaxes []float64
	var totalSalesTax float64

	for item, taxRate:= range  record{
	salesTaxes= append(salesTaxes , calculateSalesTax(item, taxRate))
	}

	for _, salesTax := range salesTaxes{
		totalSalesTax+=salesTax
	}


	return totalSalesTax

}
