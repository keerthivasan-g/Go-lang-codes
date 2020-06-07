package main
import "fmt"
func main(){
  var n float32
  fmt.Printf("Enter a floating number: ");
  fmt.Scan(&n);
  var t int
  t=int(n)
  fmt.Printf("%d\n",t);
}
