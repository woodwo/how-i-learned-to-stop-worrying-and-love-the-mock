package calculus

func Fibonacci() func() int {
	b, a := 0, 1

	return func() int {
		f := b
		b, a = a, f+a

		return f
	}
}