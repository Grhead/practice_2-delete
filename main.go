package main

import "fmt"

func main() {
	fmt.Println(ForBranch4("first", "second", "third"))
}

func ForBranch4(a string, b string, c string) (result []string) {
	result = append(result, a)
	result = append(result, b)
	result = append(result, c)
	return result
}