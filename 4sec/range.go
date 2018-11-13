package main

import (
	"fmt"
)

var pow =[]int{1,2,3,4,56,67,30}

func main() {
	fmt.Println(pow)
	// forに利用すru rangeは sliceやmapだと2つ値を返す index value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
