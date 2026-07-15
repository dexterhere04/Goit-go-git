package main

import "fmt"

func main() {
	var arr [5]byte
	fmt.Printf("%T\n", arr)

	slice := make([]byte, 5)
	fmt.Printf("%T\n", slice)
}
