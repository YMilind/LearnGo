package userlib

import ("fmt"
 "time"      
)

func reverse (s string) string {
//Rune is a Type. It occupies 32bit and is meant to represent a Unicode CodePoint. As an analogy the english characters set encoded in 'ASCII' has 128 code points
chars :=[]rune(s)

for i ,j := 0, len(chars) - 1; i < j ; i,j = i+1 ,j-1 {

chars[i] , chars[j] = chars[j] ,chars[i]

}

return string(chars)

}

func reverseString() {

var  s string
fmt.Printf("Enter the string to be reversed..")

fmt.Scanf("%s",&s)
time.Sleep(100)
fmt.Printf("Reversed String is %s\n" , reverse(s))

}


