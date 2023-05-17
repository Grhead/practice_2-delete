package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println(ForBranch1(4, 5))
	fmt.Println(ForBranch2(4.5))
	fmt.Println(ForBranch3())
}

func ForBranch1(a int, b int) int {
	return a ^ b - a*2
}

func ForBranch2(a float32) float32 {
	return a / float32(time.Now().Hour())
}

func ForBranch3() string {
	return fmt.Sprintf("%f", rand.Float64())
}

func ForBranch4(a string, b string, c string) (result []string) {
	result = append(result, a)
	result = append(result, b)
	result = append(result, c)
	return result
}