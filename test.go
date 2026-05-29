package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s int = 1
	x := strconv.FormatUint(12, 2)
	fmt.Println("conversion of 12 to binary", x)
	fmt.Println("asd", s)
}
