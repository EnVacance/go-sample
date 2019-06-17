package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int { // 複数の引数の場合は型を省略して書ける
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Brand"))
	fmt.Println(getSum(3, 4))
}
