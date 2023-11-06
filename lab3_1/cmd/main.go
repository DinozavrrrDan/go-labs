package main

import (
	"fmt"
	"lab3_1/pkg/calculate"
)

func main() {
	fmt.Print(calculate.Calculate(10, true))
	fmt.Print(calculate.Calculate(-1, true))
	fmt.Print(calculate.Calculate(10000000000000009, true))
}
