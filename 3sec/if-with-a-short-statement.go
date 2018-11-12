package main

import (
	"fmt"
	"math"
)

func pow(x,n,lim float64)float64 {
	if v := math.Pow(x,n);v <lim {
		return v
	}
	return lim
}


func main() {
	fmt.Println(
		pow(2,3,30),
		pow(3,1,2),
	)
}
