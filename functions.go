package main

import "fmt"

func ageForAlcohol(x, y int) int {
	return x * y
}

func main() {
	var responsibleAge int
	responsibleAge = ageForAlcohol(6, 3)
	fmt.Println(responsibleAge)
}
