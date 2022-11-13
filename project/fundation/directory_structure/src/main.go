package main

import (
	"fmt"
	"mymath"
)

func main() {
	num1 := 10
	num2 := 20
	var res1 = mymath.Add(num1, num2)
	var res2 = mymath.Sub(num2, num1)
	fmt.Println(res1)
	fmt.Println(res2)
}
