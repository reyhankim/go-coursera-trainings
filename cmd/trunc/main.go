package main

import (
	"fmt"	
)

func main() {
	var input float64
	fmt.Scanf("%f", &input)
	out := truncateFloat(input)
	fmt.Printf("%d\n", out)
}

func truncateFloat(input float64) int {
	truncated := int64(input)
	return int(truncated)
}