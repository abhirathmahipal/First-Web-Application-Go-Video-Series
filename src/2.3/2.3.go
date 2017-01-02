package main

import "fmt"

var (
  x int
  px *int
)

func main() {
  
  fmt.Println("Pointer before assignment")
  fmt.Println(px) 	
  x = 15
  px = &x

  fmt.Println(x, px)
  fmt.Println(*px)

}
