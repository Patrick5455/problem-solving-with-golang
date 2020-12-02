package main

/*
In this activity,
we are going to create a medical form for a doctor's office
to capture a patient's name, age, and whether they have a peanut allergy:
*/

import (
	"fmt"
)


func main() {
	fName:="Patrick"
	familyName:="Ojunde"
	age:=30
	peanutAllergy :=false

	fmt.Printf("%s\n%s\n%d\n%v", fName, familyName, age, peanutAllergy)
}



