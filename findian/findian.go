package main
import "strings"
import "fmt"
func main(){
  var s string
  fmt.Printf("Enter a string: ")
  fmt.Scan(&s)
  if ( (strings.HasPrefix(s,"i")||(strings.HasPrefix(s,"I"))) && (strings.HasSuffix(s,"n")||strings.HasSuffix(s,"N")) && (strings.ContainsAny(s,"a")||strings.ContainsAny(s,"A")) ) {
    fmt.Printf("Found\n")
  } else{
    fmt.Printf("Not Found\n")
  }
}
