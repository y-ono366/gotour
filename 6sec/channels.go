package main

import (
	"fmt"
)

func main() {
	s := []int{1,2,3,4,5,6,7,8,9,10}

	c := make(chan int)
	fmt.Println(c)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func sum(s []int, c chan int) {
	fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
