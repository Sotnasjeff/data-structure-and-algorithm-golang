package main

import (
	"fmt"
)

func main() {
	fishes := []string{"dory", "marlin", "squirtle", "gill", "nemo", "blastoise", "hank"}
	findingNemo(fishes)
}

func findingNemo(names []string) {
	for _, name := range names {
		if name == "nemo" {
			fmt.Println("Found Nemo!")
			//you can improve the code putting a break after satisfying the condition or doing something, but we need always to think
			//in worst scenarios, so this function in best scenario will be O(1) in case if the word "nemo" is the first element of array
			//in worst scenario this function will be O(n)
		}
	}
}

//this is an example in golang for O(n) or well known as Linear time it means the runtime grows proportionaly to the amount of inputs
//the algorithm receives
//Remind: Big O DOES NOT measure in seconds, the focus is how quickly our runtime grows while executing the algortihm
//[KEY WORD: Performance]
