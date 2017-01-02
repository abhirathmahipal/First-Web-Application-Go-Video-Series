package main

import "fmt"


// var can be used to declare more than one variable

var (
  x int
  s string
  i int64
)

const pi = 3.14

const (
  one = 1
  two = 2
)

func main() {

  fmt.Println("Our int variable:", x)
  x = 10
  fmt.Println("Our int variable after assignment:", x)

}
