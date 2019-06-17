package main

import "fmt"

func main() {
	var name string = "John Nash"
	// name := "John Nash" // 短い書き方 without var statement
	var age int32 = 20
	const isCool = true // like swift let, cannot re define

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool) // Typeをprint

}
