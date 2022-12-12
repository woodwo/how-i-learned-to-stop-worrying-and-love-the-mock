package main

import (
	"fmt"
	"strings"

	calculus "github.com/woodwo/calculus_lib"
)


func main () {
    f := calculus.Fibonacci()
	fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
}
