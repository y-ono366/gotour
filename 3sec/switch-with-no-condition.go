package main

import(
	"fmt"
	"time"
)


func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("午前中")
	case t.Hour() < 17:
		fmt.Println("午後")
	case t.Hour() < 21:
		fmt.Println("よる")
	default:
		fmt.Println("よーる")
	}
}
