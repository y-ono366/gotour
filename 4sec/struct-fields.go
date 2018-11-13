package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
	j int
}

func main() {
	v := Vertex{1, 2, 3}
	fmt.Println(v)
	v.X = 4
	v.j = 9
	fmt.Println(v)
}
