package main

import "fmt"

func main() {
	value := []int{2, 7, 11, 15}
	fmt.Println(twoSum(value, 9))
}

func printHowManyHelloWorldYouWantInArray(number int) []string {
	var result []string
	for i := 0; i < number; i++ {
		result = append(result, "Hello World ")
	}
	return result
}

func twoSum(nums []int, target int) []int {
	var result []int
	for index, j := range nums {
		for another, l := range nums {
			sum := j + l
			if sum == target {
				result = append(result, index, another)
			}
		}
	}
	return result
}

//this is a good example of O(n) space complexity, since in the function printHowManyHelloWorldYouWantInArray we create an empty array/slice
//and after it, we have i := 0 appending any times from input.
//remember what causes a space complexity:
// - Variables
// - Data Structures
// - Function Call
// - Allocations
