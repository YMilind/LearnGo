package main

import "fmt"

func main() {
	/* local variable definition */
	var grade string = "B"
	var marks int = 90
	fmt.Println("Enter marks to know grade")
	fmt.Scanf("%d", &marks)
	if marks < 80 && marks > 40 {
		marks = 3
	} else if marks < 90 && marks > 80 {
		marks = 2
	} else if marks > 90 {
		marks = 1
	}
	switch marks {
	case 1:
		grade = "A"
	case 2:
		grade = "B"
	case 3:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)
}
