package main

import (
	"fmt"
)

type Z interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i Z

	fmt.Println(i)
	i = &T{"Hello"}
	fmt.Println(i)
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i Z) {
	fmt.Printf("(%v, %T)\n", i, i)
}
