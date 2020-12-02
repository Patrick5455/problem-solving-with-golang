package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/*Bubble Sort

In this activity, we'll sort a given slice of numbers by swapping the values. This technique of sorting is known as the bubble sort technique. Go has built-in sorting algorithms in the sort package but I don't want you to use them; we want you to use the logic and loops you've just learned.

Steps:

    Define a slice with unsorted numbers in it.
    Print this slice to the console.
    Sort the values using swapping.
    Once done, print the now sorted numbers to the console.

Tips:

    You can do an in-place swap in Go using:

    nums[i], nums[i-1] = nums[i-1], nums[i]

    You can create a new slice using:

    var nums2 []int

    You can add to the end of a slice using:

    nums2 = append(nums2, 1)

    The following is the expected output:

    Before: [5, 8, 2, 4, 0, 1, 3, 7, 9, 6]
    After : [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]*/

func main() {

	//bring all defined functions together
	random:=10
	var bread []int

	if err:=validateInput(random); err!=nil{
		fmt.Println(err)
	} else {

		bread=getBread(random)
	}

	printBread(bread)

}

func getBread(random int)[]int{

	// create an empty slice to store digits
	var breads []int
	//randomly get digits to be added
		for i := 0; i < random; i++ {
			loaf := rand.Intn(i +random* random)
			breads = append(breads, loaf)
		}
		return breads
}

func validateInput(random int) error {
	if random < 10 || random > 100 {
		return errors.New(
			"random number can only be between 10 and 100 inclusive")
	} else {
		return nil
	}
}

func bubbleSort(bread []int)[]int{

	//var sortedBread []int
	for slice:=0; slice<len(bread); slice++{
		//fmt.Println(bread[slice])
		//fmt.Println(bread)

		if slice==0{
			continue
		}else {
			for idx, _ :=range bread[:slice+1]{

				if bread[idx] > bread[slice]{
					bread[slice], bread[idx] = bread[idx], bread[slice]
				}
			}
		}
	}
	return bread
}

func printBread (bread[]int){
	fmt.Printf("Before: %v\n", bread)
	sortedBread:= bubbleSort(bread)
	fmt.Printf("After: %v\n", sortedBread)
}