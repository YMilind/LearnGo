package main 

import "strings"
import "fmt"
func main() {
  var v int 
  var text string
var substring string
  text  = "Hello"
  v = strings.Compare("Hello", "Universe");
substring = "ello"
fmt.Scanf("%d",&v)
  fmt.Println(v);
v= strings.Index(text,substring)
fmt.Println(v)
text  =  strings.Repeat(text,1);
 fmt.Println(text)
}
