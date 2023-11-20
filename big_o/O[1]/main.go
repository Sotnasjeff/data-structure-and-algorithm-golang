package main

import (
	"fmt"
	"math/rand"
)

func main() {
	boxes := rand.Perm(100000)
	printFirstTwoBoxes(boxes)
}

// this is a O(1)
func printFirstBox(boxes []int) {
	fmt.Println(boxes[0])
}

// this is a O(2) because we've increased the number of operations and even if we increase our input, the runtime won't change, it will
// continue as constant
func printFirstTwoBoxes(boxes []int) {
	fmt.Println(boxes[0])
	fmt.Println(boxes[1])
}

//this is an example in golang for O(1) or well known as Constant time it means the runtime always will be flat independently from
//the amount of inputs your algorithm receives

//Remind: Big O DOES NOT measure in seconds, the focus is how quickly our runtime grows while executing the algortihm
//[KEY WORD: Performance]
