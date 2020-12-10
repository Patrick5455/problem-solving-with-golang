package main

import (
	"errors"
	"fmt"
)

const(
	goodScore = 450
)

func main() {

	creditScore := 500
	monthlyIncome := 1000
	laonAmount := 1000
	loanTerm := 24
	interestRate := 15

	_, errorS :=validate(creditScore, float64(monthlyIncome),
		float64(laonAmount), loanTerm)

	if errorS==nil{

		getResult(creditScore, float64(monthlyIncome), float64(laonAmount), loanTerm, float64(interestRate), 1)
	}


}


func getMonthlyPayment(loanAMount float64, interestRate float64,
	loanTerm int) float64 {


	 interestPayment := (loanAMount * interestRate)/float64(loanTerm)


	 return loanAMount/float64(loanTerm) + interestPayment
}


func validate(creditScore int, monthlyIncome float64,
	loanAmount float64, loanTerm int ) (bool, error){

	isValidationSuccessful:= creditScore>0 && monthlyIncome>0 && loanAmount>0 && loanTerm>0


	var errorS error
	if !isValidationSuccessful || loanTerm%12!=0 {
		errorS =errors.New("Validation was not Successful ")
		isValidationSuccessful = false
	}

	return isValidationSuccessful, errorS

}

func getLoanApproval(creditScore int, interestRate float64,
	monthlyIncome float64, loanTerm int, loanAmount float64)bool {

	var isApproved bool

	monthlyPayment:= getMonthlyPayment(loanAmount, interestRate,
		loanTerm)

	isMonthlyPayGr8erThanMonthlyIncome:= monthlyPayment <= monthlyIncome*0.20

	isGoodScore :=  creditScore >= goodScore && interestRate == 0.15 && isMonthlyPayGr8erThanMonthlyIncome

	if isGoodScore{isApproved = true}


	return isApproved

}


func getResult(creditScore int, monthlyIncome float64,
	loanAmount float64, loanTerm int,
	 interestRate float64, applicantNo int){


	monthlyPayment:=getMonthlyPayment(loanAmount, interestRate, loanTerm)

	isApproved:= getLoanApproval(creditScore, interestRate, monthlyIncome, loanTerm, loanAmount)


	fmt.Printf("Applicant %d \n---------\n", applicantNo)
	fmt.Println("Credit Score  : ", creditScore)
	fmt.Println("Income  : ", monthlyIncome)
	fmt.Println("Loan Amount  : ", loanAmount)
	fmt.Println("Loan Amount  : ", loanTerm)
	fmt.Println("Monthly Payment  : ", monthlyPayment)
	fmt.Println("Rate  : ", interestRate)
	fmt.Println("Total Cost  : ")
	fmt.Println("Approved  : ", isApproved)

}