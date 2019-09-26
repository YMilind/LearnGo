package main

import ("fmt"
       )

func reverse (s string) string {

chars :=[]rune(s)

for i ,j := 0, len(chars) - 1; i < j ; i,j = i+1 ,j-1 {

chars[i] , chars[j] = chars[j] ,chars[i]

}

return string(chars)

}

func main() {

var  s string
fmt.Printf("Enter the string to be reversed..")

fmt.Scanf("%s",&s)

fmt.Printf("%v\n" , reverse(s))

}


