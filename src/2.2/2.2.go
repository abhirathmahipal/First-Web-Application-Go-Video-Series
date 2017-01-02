package main

import "fmt"

func main() {
	var b1, b2 bool
	b1 = false
	b2 = false
	fmt.Println(b1, b2)

	var f1 float32
	f1 = 2.34
	fmt.Println(f1)

	var i1 int
	var i2 int64
	var i3 uint32
	fmt.Println(i1, i2, i3)

	// rune is an alias for int32
	// Unicode characters are normally stored in runes
	var r rune
	r = '('
	fmt.Println(r)

	b1 = 5 < 10
	fmt.Println(b1)

	// strings are also utf-8 encoded
	var s string
	s = "hello"
	fmt.Println(s)
	
}

