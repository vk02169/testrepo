package main

import "fmt"

func printchars() {

	var str string = "122"

	for _, c := range str {

		fmt.Printf("%d\n", c)

	}

	//	for i := 0; i < 122; i++ {

	//		fmt.Printf("%v\t%#U\n", i, i)

	//	}

}
