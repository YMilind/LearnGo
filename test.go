package main 

import "strings"
import "fmt"
func main() {
  var v int 
  var v1 string
  v1 = "Hello"
  v = strings.Compare("Hello", "World!");
  fmt.Println(v);
v1 =  strings.Repeat(v1,15);
 fmt.Println(v1)
}
