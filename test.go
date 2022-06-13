package main

import (
	"fmt"
	"time"
)

func checking() {
	time.Sleep(2 * time.Second)
	fmt.Println("Welcome")
}
func main() {
	fmt.Println("Hello")
	go checking()
	fmt.Println("bye")
	time.Sleep(3 * time.Second)
}
