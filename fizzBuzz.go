package main

import "fmt"

func main() {

	fizzBuzz()
}

func fizzBuzz(){

	for idx:=1; idx<=100; idx++{
		if idx %3 == 0{
			fmt.Println("Fizz")
		} else if idx % 5 == 0{
			fmt.Println("Buzz")
		} else if idx % 3 == 0 && idx % 5 == 0{
			fmt.Println("FizzBuzz")
		} else {
			fmt.Printf("%d\n", idx)
		}
	}
}