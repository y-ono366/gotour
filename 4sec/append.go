package main

import (
	"fmt"
)


func main() {
	// s := make([]int,6)
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 0,4,4,4)
	printSlice(s)
}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
