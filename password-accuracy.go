package main

import (
	"fmt"
	"unicode"
)

func main() {

	goodBadPassWord("This!I5A")
	goodBadPassWord(" ")
	
}


func passWordChecker(pwd string) bool{

	// we use rune type to provide support for utf-8 character
	//(multibyte character e.g chinese texts)
	pwr:= []rune(pwd)

	hasUpper:=false
	hasLower:=false
	hasSymbol:=false
	hasNUmber:=false

	for _, v := range pwr{
		if unicode.IsUpper(v){
			hasUpper = true
		}
		if unicode.IsLower(v){
			hasLower =true
		}
		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			hasSymbol =true
		}
		if unicode.IsNumber(v) || unicode.IsDigit(v){
			hasNUmber = true
		}
	}

	return hasUpper && hasLower && hasSymbol && hasNUmber
}

func goodBadPassWord(pwd string){

	//passWordStrength := false
	passWordStrength := passWordChecker(pwd)

	if passWordStrength {
		fmt.Println("Good Password")
	} else {
		fmt.Println("Bad Password")
	}

}
