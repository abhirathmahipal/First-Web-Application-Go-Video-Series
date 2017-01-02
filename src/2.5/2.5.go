package main

import "fmt"

func main() {
  a := [2][2]int {
	 {1, 1},
	 {2, 2},
  }

  b := [2][2]int {
	  {1, 1},
	  {2, 2},
  }

  c := [2][2]int {
	  {a[0][0] + b[0][0], a[0][1] + b[0][1]},
	  {a[1][0] + b[1][0], a[1][1] + b[1][1]},
  }
       fmt.Println(a)
       fmt.Println(c)
}
