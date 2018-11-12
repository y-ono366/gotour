package main

import(
	"fmt"
	"time"
)


func main() {
	fmt.Println("when's Wednesday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("明日")
	case today + 2:
		fmt.Println("明後日")
	default:
		fmt.Println("しらねぇ")
	}
}
