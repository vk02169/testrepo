package main

import "fmt"

var x int = 34
var s string = "I'm a string"
var b bool = true

func printVars() {

	str := fmt.Sprintf("%d %v %v\n", x, s, b)
	fmt.Println(str)

}
