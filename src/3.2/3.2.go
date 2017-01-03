package main

import (
	"fmt"
	"strings"
)


func main() {

   var direction string
   fmt.Scanf("%s", &direction)

   switch direction = strings.ToLower(direction); direction {
   case "n":
		fmt.Println("Go North")
   case "s":
		fmt.Println("Go South")
   case "w":
		fmt.Println("Go West")
   case "e":
		fmt.Println("Go East")

   default:
		fmt.Println("Please enter n, s, w or e")
   }

   var x int
   fmt.Scanf("%d", &x)

   switch { // switch without a variable. Using if like statements
   	
	case x >= 30:
			fmt.Println("Old")
			break
	case x >= 18 && x < 30:
			fmt.Println("Young")
			break
	case x < 18:
			fmt.Println("Kid")
			fallthrough
	default:
			fmt.Println("Sorry, you are too young")
	}
}
