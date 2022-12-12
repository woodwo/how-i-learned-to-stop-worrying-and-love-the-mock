package main

import (
	"fmt"
	"strings"
)

func fibonacci() func() int {
	b, a := 0, 1

	return func() int {
		f := b
		b, a = a, f+a

		return f
	}
}

func main () {
    f := fibonacci()
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
    fmt.Println(strings.Repeat("🐈", f()))
	fmt.Println(strings.Repeat("🐈", f()))
}
