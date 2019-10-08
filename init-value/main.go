package main

import (
	"fmt"
	"os"
)

var (
	HOME = os.Getenv("HOME")
)

func init() {
	fmt.Println(HOME)
}

func main() {
	fmt.Println(HOME)
}


