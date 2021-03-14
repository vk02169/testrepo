package main

import "fmt"

type sometype int

func typedef() {

	var x sometype

	x = 10

	fmt.Printf("type of x = %T\n", x)

}
