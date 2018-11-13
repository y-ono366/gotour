package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printLenCap("a", a)

	s := make([]int, 0, 5)
	printLenCap("s", s)
}

func printLenCap(n string,s []int){
	fmt.Printf("%s len=%d cap=%d %v\n",
		n, len(s), cap(s), s)
}
