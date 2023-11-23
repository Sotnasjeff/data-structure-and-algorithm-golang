package main

import "fmt"

func main() {
	boxes := []string{"a", "b", "c", "d", "e"}
	logAllPairs(boxes)
}

func logAllPairs(array []string) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("%s %s\n", array[i], array[j])
		}
	}
}

//this is an example in golang for O(n^2) or called quadratic time, what it means the runtime will increase quadraticly who it's horrible
//when there are some nested loops, you can make sure the algorithm will not be good in performance.
