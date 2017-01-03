package main

import "fmt"

func main() {
	x := map[int]int{1: 10, 2: 20}
	fmt.Println("value for 3:", x[3])


	_, ok := x[1]

	if ok {
		fmt.Println("Value of 1 exists")
	}

	_, ok2 := x[100]

	if ok2 {
		fmt.Println("Value for 100 exists")
	} else {
		fmt.Println("Value for 100 does not exist")
	}

}
