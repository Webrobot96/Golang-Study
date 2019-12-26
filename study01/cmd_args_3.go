package main

import (
	"fmt"
	"os"
	"string"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
