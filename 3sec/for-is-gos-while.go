package main

import (
	"fmt"
)

func main() {
	var sum int = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
