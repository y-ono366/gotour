package main

import(
	"fmt"
)


func main() {
	ohno("ohno ")
	defer fmt.Printf("wide ")
	fmt.Printf("world ")

}

func ohno(str string) {
	defer fmt.Printf("is ")
	fmt.Printf(str)
}
