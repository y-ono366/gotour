package main

import (
	"fmt"
)
type Vertx struct {
	Lit,Mit float64
}


var m map[string]Vertx

func main() {
	m = make(map[string]Vertx)
	m["test"] = Vertx{100,100}
	m["test2"] = Vertx{100,100}
	fmt.Println(m)
}
