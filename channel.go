package main

import "fmt"

func main() {
	var length int
var name string
	fmt.Println("Enter the queue length")
	fmt.Scanf("%d", &length)
	queue := make(chan string, length)
	for q := 0; q < length; q++ {
		fmt.Println("Enter names")
		fmt.Scanf("%s", &name)
		queue <- name

	}
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
