package main

import "fmt"
import "math/rand"

func main() {
	fmt.Println(ForBranch3())
}

func ForBranch3() string {
	return fmt.Sprintf("%f", rand.Float64())
}

