package main

import (
	"fmt"
	"packages/calculus"
	"strings"
)

func main () {
    f := calculus.Fibonacci()
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
}
