package main

import (
	"fmt"
	"os"
)

func init() {
	/*;)*/ cnt, err := os.ReadFile("test.asm")
	if err != nil {
		fmt.Println("OH FUCK")
	}
	fmt.Println(cnt)
}

func main() {

}
