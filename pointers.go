package main

import (
	"fmt"
)

/*
In this activity, your job is to finish some code a co-worker started. Here, we have some unfinished code for you to complete. Your task is to fill in the missing code, where the comments are to swap the values of a and b. The swap function only accepts pointers and doesn't return anything:

package main
import "fmt"
func main() {
  a, b := 5, 10
  // call swap here
  fmt.Println(a == 10, b == 5)
}
func swap(a *int, b *int) {
  // swap the values here
}

    Call the swap function, ensuring you are passing a pointer.
    In the swap function, assign the values to the other pointer, ensuring you dereference the values.

    The following is the expected output:

    true true
*/

func main() {
	a, b := 5, 10
	// call swap here
	//NB: We cannot pass the literal values of variables a and b to swap.
	//This is because swap collects pointer arguments. So to pass in the argument type,
	//we can just use the & symbol with the variable
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}


func swap(a *int, b *int) {

	//assign the value of pointer a (dereferenced to a temporary variable)
	temp:= *a
	//assign the value of pointer b to value of a (both dereference)
	*a = *b
	// assign temp (a value type) to pointer b (dereference)
	*b = temp
}
