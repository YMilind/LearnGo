package main

import "fmt"

func main() {
	var name [5]string
	iter := 4
	fmt.Println("Enter 5 words")
	for x := 0; x <= iter; x++ {
		fmt.Scanf("%s", &name[iter])
	}
	fmt.Println("Reverese order is ")
	fmt.Println( swap(&name[0],&name[1],&name[2],&name[3],&name[4]))

}

func swap(a, b, c, d, e *string) (string, string, string, string, string) {
	return *e, *d, *c, *b, *a
}
