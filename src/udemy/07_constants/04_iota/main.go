package main

import "fmt"

const (
	_ = iota
  B = iota * 10
  C = iota * 10
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}
