package main

import (
	"fmt"
)

func main() {
	fmt.Println(ForBranch1(4, 5))
}

func ForBranch1(a int, b int) int {
	return a ^ b - a*2
}
