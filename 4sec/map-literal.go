package main

import (
	"fmt"
)
type Vertx struct {
	Lit,Mit float64
}

var m = map[string]Vertx {
	"fuck in ":Vertx{
		300,400,
	},
	"world":Vertx{
		400,500,
	},
}

func main() {
	fmt.Println(m)
}
