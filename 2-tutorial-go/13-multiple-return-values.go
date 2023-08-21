// Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
