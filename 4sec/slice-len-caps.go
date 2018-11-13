package main

import (
	"fmt"
)

func main() {
	s := []int{2,3,4,5,65,100,30,20}
	printLenCap(s)
	printLenCap(s[4:])
	printLenCap(s[2:8])
}

func printLenCap(s []int){
	fmt.Printf("len:%g  cap:%g  %v\n",len(s),cap(s),s)
}
