package assert

func Sub(args ...int) int {
	a := args[0]
	for _, b := range args[1:] {
		a -= b
	}
	return a
}
