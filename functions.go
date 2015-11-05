package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7

}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := vals()
	fmt.Println(c)
	res := plus(1, 2)
	fmt.Println("1+2=", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
