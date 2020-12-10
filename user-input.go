package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var userCollection map[int]string

const exactLength = 2

func main() {

	// to run the program, find below example format of what to type in console:
	// go run <name_of_file>.go 305

	userCollection=make(map[int]string)

	userCollection[305]="Sue"
	userCollection[204]="Bob"
	userCollection[631]="Jake"
	userCollection[073]="Tracy"

	isValid, err := validate()

	if isValid{
		key, _:=strconv.Atoi(os.Args[1])
		fmt.Println("Key: ", key)
		fmt.Printf("value: %v\n",userCollection[key])
	}else {
		fmt.Print("Error-> ", err)
	}

}

func validate()(bool,error){

	var err error
	var isValid bool
	var isKey bool

	for k, _ := range userCollection{
		if test:=strconv.Itoa(k); test == os.Args[1]{
			fmt.Println("Yay! ---- The key exists!")
			isKey = true
		}
	}

	if len(os.Args) > exactLength || len(os.Args) < exactLength || isKey==false{
		err= errors.New("It is either you are passing more than one key or that key is not in our collection\n")
	} else {
		err = nil
	}

	isValid = err==nil && isKey
	fmt.Println("isvalid:",isValid)

	return isValid, err
}



