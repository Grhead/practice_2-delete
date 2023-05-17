package main

import "fmt"
import "time"
import "math/rand"

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

func ForBranch3() string {
	return fmt.Sprintf("%f", rand.Float64())
}




func ForBranch4() []string {

}