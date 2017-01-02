package main

import "fmt"

func main() {
  // Internally slices work with pointers
  s := []int{1, 2, 3}
  fmt.Println(s)
  fmt.Println(s[0])

  var s1 []int
  s1 = make([]int, 3, 15) // 3rd 15 the max capacity
  fmt.Println(s1)
  fmt.Println("Length is", len(s1), "Cap is", cap(s1))
  
  // appending data increases the length
  s1 = append(s1, 1)
  fmt.Println("Slice after appending 1", s1)

  // copy - will copy as much as possible
  // i.e if 20 is copied into 15, the first 15 elements
  // will be copied from the 20 size slice to the
  // 15 element slice

  // Arrays are faster but more rigid

  a := [3]int{1,2,3}
  fmt.Println("Printing array", a)


}
