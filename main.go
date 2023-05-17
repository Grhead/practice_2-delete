package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(ForBranch1(4, 5))
	fmt.Println(ForBranch2(4.5))
}

func ForBranch1(a int, b int) int {
	return a ^ b - a*2
}

func ForBranch2(a float32) float32 {
	return a / float32(time.Now().Hour())

}
