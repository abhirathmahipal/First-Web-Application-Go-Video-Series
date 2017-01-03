package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	var sum int

	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	fmt.Println(sum)

	var y int

	for y < 10 {
		fmt.Scanf("%d", &y)
	}
	fmt.Println("Finally! Greater than 10")

//	Infinite loops
//	for {
//
//	}


}
