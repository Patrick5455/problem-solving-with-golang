package main

import "fmt"

/*Loop over map data and print out the word with the highest count
"Gonna": 3,
  "You": 3,
  "Give": 2,
  "Never": 1,
  "Up":  4,
*/

var words = map[string]int{
	"Gonna": 3,
	"You": 3,
	"Give": 2,
	"Never": 1,
	"Up":  4,
}

func main() {

	popular(words)

}



func popular(words map[string]int){

	max, popularWord := 0, " "

	for k, v := range words{
		if v>max{
			popularWord = k
			max = v
		}}
	fmt.Printf("Most Popular word : %s \nWith a count of : %d", popularWord, max)

}
