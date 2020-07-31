package main

import (
	"fmt"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	var total int = Sum(5,  5)
	fmt.Println(total)
}
