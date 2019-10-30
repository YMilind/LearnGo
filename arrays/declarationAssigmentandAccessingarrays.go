package main

import "fmt"

func main() {

	var n [10]int
	// n above is an array of 10 intergers
	var i, j int

	//initialize arrays

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	//display arrays
	for j = 0; j < 10; j++ {
		fmt.Printf("Element %d is %d..\n", j, n[j])
	}
}
