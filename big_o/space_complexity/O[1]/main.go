package main

import "fmt"

func main() {
	printSomethingHowManyTimesYouWant(10)
}

func printSomethingHowManyTimesYouWant(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Hello World!!")
	}
}

//this is an example of space complexity using Big O notation, in this case, We've only stored only 1 variable in memory whose it's: "i := 0"
//no matter the input, but what it's inside of this function.

//this is a case of constant, so the big O it's O(1)
