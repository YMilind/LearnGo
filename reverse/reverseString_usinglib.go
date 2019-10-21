package main

import (
	"fmt"
	"userlib"
)

func main() {
	var s string
	fmt.Println("Enter the string to be reversed")

	fmt.Scanf("%s", &s)

	//userlib.reverseStringlib.reverse(s)
	//userlib.Reverse(s)
	fmt.Println(userlib.Reverse(s))
}
