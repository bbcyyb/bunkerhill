package main

import (
	"fmt"
)

var ()

func main() {
	fmt.Println("Hello, playground")
	log := NewLogger(10000)

	log.Debug("123456789")
}
