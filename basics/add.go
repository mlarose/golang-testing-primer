package basics

// Add computes the sum of variadic list of integers and returns it
func Add(ops ...int) int {
	accumulator := 0
	for _, op := range ops {
		accumulator += op
	}
	return accumulator
}
